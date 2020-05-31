package sdrive

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"streetlity-cdn/config"
	"strings"
	"time"
)

type UploadType int

const (
	None UploadType = iota
	Override
	Rename
)

func UploadFile(path string, data []byte, utype UploadType) (p string, e error) {
	reg := regexp.MustCompile("^[\\w\\-. ]+$")
	log.Println(utype == None)
	if !reg.MatchString(path) {
		log.Println("[SDrive]", "invalid filename", path)
		return p, errors.New("Invalid filename")
	}

	prefix := GetPrefixDirectory()

	p = path
	dir := filepath.Join(config.Config.Location, prefix)
	base := filepath.Base(p)
	ext := filepath.Ext(base)

	if _, e = os.Stat(dir); os.IsNotExist(e) {
		log.Println("[SDrive]", "Create directory", dir)
		if e = os.MkdirAll(dir, os.ModeDir); e != nil {
			log.Println("[SDrive]", "Error when create directory", e.Error())
			return
		}
	}

	count := 1
	_, e = os.Stat(filepath.Join(dir, p))

	for !os.IsNotExist(e) {
		switch utype {
		case None:
			log.Println("Hi mom i'm at none")
			return p, errors.New("file is existed")
		case Override:
		case Rename:
			log.Println("Hi mom i'm at rename")
			s := strings.Split(base, ".")
			p = fmt.Sprintf("%s(%d)%s", s[0], count, ext)
			_, e = os.Stat(filepath.Join(dir, p))
			count += 1
		}
	}
	e = ioutil.WriteFile(filepath.Join(dir, p), data, 0777)

	p = filepath.Join(prefix, p)
	return
}

func DeleteFiles(path []string) (e error) {
	prefix := GetPrefixDirectory()
	for _, p := range path {
		p = filepath.Join(config.Config.Location, prefix)
		_, e = os.Stat(p)

		if os.IsNotExist(e) {
			continue
		}

		os.Remove(p)
	}

	return
}

func DownloadFile(path string) (reader io.Reader, e error) {
	path = filepath.Join(config.Config.Location, path)

	reader, e = os.Open(path)

	return
}

//Generate path for storaging
func GetPrefixDirectory() (prefix string) {
	now := time.Now()
	d := now.Format("2006-01-02")
	t := strconv.FormatInt(int64(now.Hour()), 10)

	prefix = filepath.Join(d, t)
	return
}
