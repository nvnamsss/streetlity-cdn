package sdrive

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type UploadType int

const (
	None UploadType = iota
	Override
	Rename
)

func UploadFile(path string, data []byte, utype UploadType) (p string, e error) {
	p = path
	_, e = os.Stat(p)
	count := 1
	dir := filepath.Dir(p)
	base := filepath.Base(p)
	ext := filepath.Ext(base)

	for !os.IsNotExist(e) {
		s := strings.Split(base, ".")
		name := fmt.Sprintf("%s(%d)%s", s[0], count, ext)
		p = filepath.Join(dir, name)
		_, e = os.Stat(p)
		count += 1
	}

	ioutil.WriteFile(p, data, 0777)
	return
}

func DeleteFiles(path []string) (e error) {
	for _, p := range path {
		_, e = os.Stat(p)

		if os.IsNotExist(e) {
			continue
		}

		os.Remove(p)
	}

	return
}
