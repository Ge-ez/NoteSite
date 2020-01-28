package handlers

import(
    	"io"
    	"mime/multipart"
    	"os"
    	"path/filepath"
)

func UploadFile(mf *multipart.File, name string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	path := filepath.Join(wd, "views", "assets", "files", name)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	io.Copy(file, *mf)
	return nil
}
