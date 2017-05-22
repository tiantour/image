package image

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/tiantour/conf"
)

// Server server
type Server struct{}

// NewServer new server
func NewServer() *Server {
	return &Server{}
}

// Net net upload
// date 2017-05-22
// author andy.jiang
func (s Server) Net(url string) (string, error) {
	body, err := NewFile().Net(url)
	if err != nil {
		return "", err
	}
	return s.Local(body)
}

// Local local upload
// date 2017-05-22
// author andy.jiang
func (s Server) Local(imageByte []byte) (imagePath string, err error) {
	path := fmt.Sprintf("%s/%s",
		conf.NewConf().Server.Upload,
		NewFile().Name(),
	)
	f, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	data := bytes.NewBuffer(imageByte) // io.reader
	_, err = io.Copy(f, data)          // write
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s",
		conf.NewConf().Server.Domain,
		path,
	), nil
}
