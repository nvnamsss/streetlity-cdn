package sdrive

type UploadType int

const (
	None UploadType = iota
	Override
	Rename
)

func UploadFile(path string, utype string) (p string, e error) {
	return
}

func UploadFiles(path []string, utype int) (p map[string]string, e error) {
	return
}

func DeleteFiles(path []string) (e error) {
	return
}
