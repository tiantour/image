package image

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/tiantour/fetch"
)

// File file
type File struct{}

// NewFile new file
func NewFile() *File {
	return &File{}
}

// Name get file name
func (f File) Name() string {
	now := time.Now().UnixNano()
	name := strconv.FormatInt(now, 10)
	return fmt.Sprintf("file_%s.jpg", name)
}

// Net read net file
func (f File) Net(url string) ([]byte, error) {
	return fetch.Cmd(fetch.Request{
		Method: "GET",
		URL:    url,
	})
}

// Local read local file
func (f File) Local(path string) ([]byte, error) {
	imageFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(imageFile)
}
