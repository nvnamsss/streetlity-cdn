package sdrive

import (
	"errors"
	"log"
	"os"
	"regexp"
)

func Modify(path string, data []byte) (e error) {
	reg := regexp.MustCompile("^[\\w\\-. ]+$")
	if !reg.MatchString(path) {
		log.Println("[SDrive]", "invalid filename", path)
		return errors.New("Invalid filename")
	}

	path = GetPath(path)
	if _, e = os.Stat(path); os.IsNotExist(e) {
		return
	}

	f, _ := os.Open(path)
	_, e = f.WriteAt(data, 0)
	return
}
