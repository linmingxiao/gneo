package gneo

import (
	"net/http"
	"os"
)

type onlyFilesFS struct{
	fs http.FileSystem
}

type neuteredReaddirFile struct {
	http.File
}

func Dir(root string, listDirectory bool) http.FileSystem{
	fs:= http.Dir(root)
	if listDirectory{
		return fs
	}
	return &onlyFilesFS{fs}
}


func (fs onlyFilesFS) Open(name string)	(http.File, error){
	f, err := fs.fs.Open(name)
	if err != nil{
		return nil, err
	}
	return neuteredReaddirFile{f}, nil
}

func (f neuteredReaddirFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}







