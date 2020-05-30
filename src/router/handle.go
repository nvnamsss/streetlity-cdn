package router

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func download(w http.ResponseWriter, req *http.Request) {
	reader, _ := ReadImage("AdobeStock_53119595.jpeg")
	w.Header().Set("Content-Disposition", "attachment; filename=image.png")
	w.Header().Set("Content-Type", req.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", req.Header.Get("Content-Length"))

	io.Copy(w, reader)
}

func upload(w http.ResponseWriter, req *http.Request) {
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
}

func delete(w http.ResponseWriter, req *http.Request) {

}

func Handle(router *mux.Router) {
	log.Println("[Router]", "init handle")
	router.HandleFunc("/", download).Methods("GET")
	router.HandleFunc("/", upload).Methods("POST")
	router.HandleFunc("/", delete).Methods("DELETE")
}
