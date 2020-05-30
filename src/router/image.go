package router

import (
	"io"
	"os"
)

func ReadImage(path string) (reader io.Reader, e error) {

	return os.Open(path)
}

func UploadImage(path string, data []byte) {

}
