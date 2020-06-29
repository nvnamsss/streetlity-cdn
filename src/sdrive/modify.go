package sdrive

import (
	"errors"
	"log"
	"os"
	"regexp"
)

func Modify(path string, data []byte) (e error) {
	reg := regexp.MustCompile("^(.+)\\/([^/]+)$")
	if !reg.MatchString(path) {
		log.Println("[SDrive]", "modify", "invalid filename", path)
		return errors.New("Invalid filename")
	}

	log.Println("[SDrive]", "modify", "data", string(data))
	path = GetPath(path)
	if _, e = os.Stat(path); os.IsNotExist(e) {
		log.Println("[SDrive]", "modify", "path was not exist", path)
		return
	}

	f, e := os.OpenFile(path, os.O_WRONLY, os.ModeAppend)
	if e != nil {
		log.Println("[SDrive]", "modify", "open", e.Error())
	}
	defer f.Close()

	_, e = f.Write(data)
	if e != nil {
		log.Println("[SDrive]", "modify", "write", e.Error())
	}
	return
}
