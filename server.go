package image

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/tiantour/conf"
)

// Net
func (s *server) Net(imageURL string) (imagePath string, err error) {
	imageByte, err := File.net(imageURL)
	if err != nil {
		return
	}
	imagePath, err = Server.Local(imageByte)
	return
}

// Local
func (s *server) Local(imageByte []byte) (imagePath string, err error) {
	filePath := fmt.Sprintf("%s/%s", conf.Data.Server.Upload, File.name()) // filePath
	f, err := os.Create(filePath)                                          // Create
	if err != nil {
		return
	}
	defer f.Close()
	data := bytes.NewBuffer(imageByte) // io.reader
	_, err = io.Copy(f, data)          // write
	if err != nil {
		return
	}
	imagePath = fmt.Sprintf("%s/%s", conf.Data.Server.Domain, filePath) // image path
	return
}
