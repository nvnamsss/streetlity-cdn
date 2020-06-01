package router

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"streetlity-cdn/sdrive"

	"github.com/gorilla/mux"
	"github.com/nvnamsss/goinf/pipeline"
)

func download(w http.ResponseWriter, req *http.Request) {
	req.URL.RawQuery, _ = url.QueryUnescape(req.URL.RawQuery)
	p := pipeline.NewPipeline()
	stage := pipeline.NewStage(func() (str struct {
		Filename string
	}, e error) {
		query := req.URL.Query()

		filenames, ok := query["f"]

		if !ok {
			return str, errors.New("f param is missing")
		}

		str.Filename = filenames[0]

		return
	})

	p.First = stage
	e := p.Run()

	if e != nil {
		var res Response = Response{}
		res.Error(e)
		WriteJson(w, res)
		return
	}

	filename := p.GetStringFirstOrDefault("Filename")

	reader, e := sdrive.DownloadFile(filename)

	if e != nil {
		var res Response = Response{}
		res.Error(e)
		WriteJson(w, res)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(filename))
	w.Header().Set("Content-Type", req.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", req.Header.Get("Content-Length"))

	io.Copy(w, reader)
}

func upload(w http.ResponseWriter, req *http.Request) {
	var res struct {
		Response
		Paths map[string]Response
	}
	res.Status = true
	res.Paths = make(map[string]Response)

	req.ParseForm()
	p := pipeline.NewPipeline()
	stage := pipeline.NewStage(func() (str struct {
		Filename   []string
		UploadType int
	}, e error) {
		query := req.URL.Query()
		files, ok := query["f"]

		if !ok {
			return str, errors.New("f param is missing")
		}

		t, ok := query["utype"]

		if ok {
			if v, e := strconv.Atoi(t[0]); e == nil {
				str.UploadType = v
				log.Println(v)
			}
		}

		str.Filename = files
		return
	})
	p.First = stage
	res.Error(p.Run())

	if !res.Status {
		WriteJson(w, res)
		return
	}

	filenames := p.GetString("Filename")
	utype := p.GetIntFirstOrDefault("UploadType")

	log.Println(filenames)
	// res.Paths = filemap
	req.ParseMultipartForm(32 << 20) // limit your max input length!

	for _, f := range filenames {
		file, _, e := req.FormFile(f)
		if e != nil {
			res.Paths[f] = Response{Status: false, Message: "cannot find this file in the form"}
			continue
		}

		defer file.Close()

		var buf bytes.Buffer
		io.Copy(&buf, file)
		p, e := sdrive.UploadFile(f, buf.Bytes(), sdrive.UploadType(utype))
		if e != nil {
			res.Paths[f] = Response{Status: false, Message: e.Error()}
		} else {
			res.Paths[f] = Response{Status: true, Message: p}
		}

		buf.Reset()
	}

	WriteJson(w, res)
}

func delete(w http.ResponseWriter, req *http.Request) {
	var res Response = Response{Status: true}

	req.URL.RawQuery, _ = url.QueryUnescape(req.URL.RawQuery)
	p := pipeline.NewPipeline()
	stage := pipeline.NewStage(func() (str struct {
		Filename []string
	}, e error) {
		query := req.URL.Query()

		files, ok := query["f"]
		if !ok {
			return str, errors.New("f param is missing")
		}

		str.Filename = files

		return
	})
	p.First = stage
	res.Error(p.Run())

	if res.Status {
		filenames := p.GetString("Filename")
		sdrive.DeleteFiles(filenames)
	}

	WriteJson(w, res)
}

func Handle(router *mux.Router) {
	log.Println("[Router]", "init handle")
	router.HandleFunc("/", download).Methods("GET")
	router.HandleFunc("/", upload).Methods("POST")
	router.HandleFunc("/", delete).Methods("DELETE")
}
