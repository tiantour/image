package image

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// Server server
type Server struct{}

// NewServer new server
func NewServer() *Server {
	return &Server{}
}

// Local local upload
func (s *Server) Local(args *Image) (string, error) {
	args.Name = NewFormat().Name(args)
	args.Path = NewFormat().Path(args)

	err := s.do(args)
	if err != nil {
		return "", err
	}

	path := fmt.Sprintf("%s/%s", cfs.Host, args.Path)
	return path, nil
}

// do
func (s *Server) do(args *Image) error {
	path := fmt.Sprintf("%s/%s", cfs.Bucket, args.Path)

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	body := bytes.NewBuffer(args.Body)
	_, err = io.Copy(f, body)
	return err
}
