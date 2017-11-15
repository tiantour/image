package image

import (
	"bytes"
	"fmt"
	"log"

	"github.com/upyun/go-sdk/upyun"
)

var up *upyun.UpYun

// Upyun upyun
type Upyun struct{}

// NewUpyun new upyun
func NewUpyun() *Upyun {
	if UpyunBucket == "" || UpyunHost == "" || UpyunUname == "" || UpyunPasswd == "" {
		log.Fatal("image conf is null")
	}
	up = upyun.NewUpYun(&upyun.UpYunConfig{
		Bucket:   UpyunBucket,
		Operator: UpyunUname,
		Password: UpyunPasswd,
	})
	return &Upyun{}
}

// Net net upload
func (u Upyun) Net(url string) (string, error) {
	body, err := NewFile().Net(url)
	if err != nil {
		return "", err
	}
	return u.Local(body)
}

// Local local upload
func (u Upyun) Local(body []byte) (string, error) {
	path := fmt.Sprintf("%s/%s",
		UpyunBucket,
		NewFile().Name(),
	)
	err := up.Put(&upyun.PutObjectConfig{
		Path:   path,
		Reader: bytes.NewBuffer(body),
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", UpyunHost, path), nil
}
