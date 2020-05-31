package sdrive

import (
	"errors"
	"fmt"
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

	if !reg.MatchString(path) {
		log.Println("[SDrive]", "invalid filename", path)
		return p, errors.New("Invalid filename")
	}

	p = filepath.Join(GetPath(), path)
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

//Generate path for storaging
func GetPath() string {
	now := time.Now()
	d := now.Format("2006-01-02")
	t := strconv.FormatInt(int64(now.Hour()), 10)

	return filepath.Join(config.Config.Location, d, t)
}
