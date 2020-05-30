package router

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/nvnamsss/goinf/pipeline"
)

func download(w http.ResponseWriter, req *http.Request) {
	p := pipeline.NewPipeline()
	stage := pipeline.NewStage(func() (str struct {
		Filename string
	}, e error) {
		query := req.URL.Query()
		filenames, ok := query["filename"]

		if !ok {
			return str, errors.New("filename param is missing")
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
	var res Response = Response{Status: true}

	req.ParseMultipartForm(32 << 20) // limit your max input length!
	var buf bytes.Buffer
	file, header, err := req.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	name := strings.Split(header.Filename, ".")
	fmt.Printf("File name %s\n", name[0])
	io.Copy(&buf, file)

	contents := buf.String()
	ioutil.WriteFile("backup.sql", buf.Bytes(), 0777)
	fmt.Println(contents)
	buf.Reset()

	WriteJson(w, res)
}

func delete(w http.ResponseWriter, req *http.Request) {

}

func Handle(router *mux.Router) {
	log.Println("[Router]", "init handle")
	router.HandleFunc("/", download).Methods("GET")
	router.HandleFunc("/", upload).Methods("POST")
	router.HandleFunc("/", delete).Methods("DELETE")
}
