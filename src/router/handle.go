package router

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
	"streetlity-cdn/sdrive"

	"github.com/gorilla/mux"
	"github.com/nvnamsss/goinf/pipeline"
)

func download(w http.ResponseWriter, req *http.Request) {
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

	reader, e := ReadImage(filename)

	if e != nil {
		var res Response = Response{}
		res.Error(e)
		WriteJson(w, res)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=image.png")
	w.Header().Set("Content-Type", req.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", req.Header.Get("Content-Length"))

	io.Copy(w, reader)
}

func upload(w http.ResponseWriter, req *http.Request) {
	var res struct {
		Response
		Paths map[string]string
	}
	res.Status = true

	req.ParseForm()
	p := pipeline.NewPipeline()
	stage := pipeline.NewStage(func() (str struct {
		Filename   []string
		UploadType int
	}, e error) {
		form := req.PostForm

		filenames, ok := form["filename"]
		if !ok {
			return str, errors.New("filename param is missing")
		}

		t, ok := form["utype"]

		if ok {
			if v, e := strconv.Atoi(t[0]); e == nil {
				str.UploadType = v
			}
		}

		str.Filename = filenames
		return
	})
	p.First = stage
	res.Error(p.Run())

	if res.Status {
		WriteJson(w, res)
		return
	}

	filenames := p.GetString("Filename")
	utype := p.GetIntFirstOrDefault("UploadType")

	filemap, e := sdrive.UploadFiles(filenames, int(utype))
	res.Error(e)

	res.Paths = filemap
	req.ParseMultipartForm(32 << 20) // limit your max input length!
	var buf bytes.Buffer
	file, header, err := req.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(&buf, file)
	log.Println(header.Filename)
	// ioutil.WriteFile(value, buf.Bytes(), 0777)
	buf.Reset()
	// for key, value := range filemap {

	// }

	WriteJson(w, res)
}

func delete(w http.ResponseWriter, req *http.Request) {
	var res Response = Response{Status: true}

	req.ParseForm()
	p := pipeline.NewPipeline()
	stage := pipeline.NewStage(func() (str struct {
		Filename []string
	}, e error) {
		form := req.PostForm

		files, ok := form["f"]
		if !ok {
			return str, errors.New("f param is missing")
		}

		str.Filename = files

		return
	})
	p.First = stage
	res.Error(p.Run())

	if res.Status {

	}

	WriteJson(w, res)
}

func Handle(router *mux.Router) {
	log.Println("[Router]", "init handle")
	router.HandleFunc("/", download).Methods("GET")
	router.HandleFunc("/", upload).Methods("POST")
	router.HandleFunc("/", delete).Methods("DELETE")
}
