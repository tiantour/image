package image

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

// Server server
type Server struct{}

// NewServer new server
func NewServer() *Server {
	if ServerBucket == "" || ServerHost == "" {
		log.Fatal("image conf is null")
	}
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
		ServerBucket,
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
	return fmt.Sprintf("%s/%s", ServerHost, path), nil
}
