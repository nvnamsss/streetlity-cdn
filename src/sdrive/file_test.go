package sdrive_test

import (
	"log"
	"path/filepath"
	"streetlity-cdn/sdrive"
	"testing"
)

func TestUploadFile(t *testing.T) {
	path, e := sdrive.UploadFile("abc.meo", []byte("abc"), sdrive.Rename)
	prefix := sdrive.GetPrefixDirectory()
	if e != nil {
		log.Println(e.Error())
	}

	if path != filepath.Join(prefix, "abc.meo") {
		t.Fatalf("UploadFile failed, expected %v got %v", filepath.Join(prefix, "abc.meo"), path)
	}
	t.Logf("Completed")
}

//TestGetPath run the sdrive.GetPath(), due to it's depend on the time and config
//thus it's recommend to run in debug mode rather than common
func TestGetPath(t *testing.T) {
	s := sdrive.GetPrefixDirectory()

	log.Println(s)

	t.Logf("Completed")
}
