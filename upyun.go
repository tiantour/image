package image

import (
	"bytes"
	"fmt"

	"github.com/upyun/go-sdk/upyun"
)

// Upyun upyun
type Upyun struct{}

// NewUpyun new upyun
func NewUpyun() *Upyun {
	return &Upyun{}
}

// Local local upload
func (u *Upyun) Local(args *File) (string, error) {
	args.Name = NewFormat().Name(args)
	args.Path = NewFormat().Path(args)

	err := u.do(args)
	if err != nil {
		return "", err
	}

	path := fmt.Sprintf("%s/%s", cfu.Host, args.Path)
	return path, nil
}

// do
func (u *Upyun) do(args *File) error {
	path := fmt.Sprintf("%s/%s", cfu.Bucket, args.Path)

	return up.Put(&upyun.PutObjectConfig{
		Path:   path,
		Reader: bytes.NewReader(args.Body),
	})
}
