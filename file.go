package image

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/tiantour/fetch"
)

// Prefix file pre fix
var Prefix string

// File file
type File struct{}

// NewFile new file
func NewFile() *File {
	return &File{}
}

// Name get file name
func (f *File) Name() string {
	if Prefix == "" {
		Prefix = "file"
	}
	return fmt.Sprintf("%s/%d.jpg", Prefix, time.Now().UnixNano())
}

// Net read net file
func (f *File) Net(url string) ([]byte, error) {
	return fetch.Cmd(fetch.Request{
		Method: "GET",
		URL:    url,
	})
}

// Local read local file
func (f *File) Local(path string) ([]byte, error) {
	imageFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(imageFile)
}
