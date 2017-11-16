package image

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/tiantour/conf"
)

var (
	cfs = conf.NewConf().Image["server"]
)

// Server server
type Server struct{}

// NewServer new server
func NewServer() *Server {
	return &Server{}
}

// Net net upload
func (s Server) Net(url string) (string, error) {
	body, err := NewFile().Net(url)
	if err != nil {
		return "", err
	}
	return s.Local(body)
}

// Local local upload
func (s Server) Local(imageByte []byte) (imagePath string, err error) {
	path := fmt.Sprintf("%s/%s",
		cfs.Bucket,
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
	return fmt.Sprintf("%s/%s", cfs.Host, path), nil
}
