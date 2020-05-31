package sdrive_test

import (
	"log"
	"streetlity-cdn/sdrive"
	"testing"
)

func TestUploadFile(t *testing.T) {
	sdrive.UploadFile("abc.meo", []byte("abc"), sdrive.Rename)

	t.Logf("Completed")
}

//TestGetPath run the sdrive.GetPath(), due to it's depend on the time and config
//thus it's recommend to run in debug mode rather than common
func TestGetPath(t *testing.T) {
	s := sdrive.GetPath()

	log.Println(s)

	t.Logf("Completed")
}
