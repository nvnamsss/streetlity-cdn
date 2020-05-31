package sdrive_test

import (
	"streetlity-cdn/sdrive"
	"testing"
)

func TestUploadFile(t *testing.T) {
	sdrive.UploadFile("abc.meo", []byte("abc"), sdrive.Rename)

	t.Logf("Completed")
}
