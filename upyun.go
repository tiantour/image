package image

import (
	"bytes"
	"fmt"

	"github.com/tiantour/conf"
	"github.com/upyun/go-sdk/upyun"
)

var (
	up  *upyun.UpYun
	cfu = conf.NewImage().Data["upyun"]
)

func init() {
	up = upyun.NewUpYun(&upyun.UpYunConfig{
		Bucket:   cfu.Bucket,
		Operator: cfu.Uname,
		Password: cfu.Passwd,
	})
}

// Upyun upyun
type Upyun struct{}

// NewUpyun new upyun
func NewUpyun() *Upyun {
	return &Upyun{}
}

// Net net upload
func (u *Upyun) Net(url string) (string, error) {
	body, err := NewFile().Net(url)
	if err != nil {
		return "", err
	}
	return u.Local(body)
}

// Local local upload
func (u *Upyun) Local(body []byte) (string, error) {
	path := fmt.Sprintf("%s/%s", cfu.Bucket, NewFile().Name())
	err := up.Put(&upyun.PutObjectConfig{
		Path:   path,
		Reader: bytes.NewBuffer(body),
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", cfu.Host, path), nil
}
