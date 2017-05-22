package image

import (
	"bytes"
	"fmt"

	"github.com/tiantour/conf"
	"github.com/upyun/go-sdk/upyun"
)

// Upyun upyun
type Upyun struct{}

// NewUpyun new upyun
func NewUpyun() *Upyun {
	return &Upyun{}
}

var up = upyun.NewUpYun(&upyun.UpYunConfig{
	Bucket:   conf.NewConf().Upyun.Bucket,
	Operator: conf.NewConf().Upyun.Username,
	Password: conf.NewConf().Upyun.Passwd,
})

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
	host := conf.NewConf().Upyun.Host
	path := fmt.Sprintf("%s/%s",
		conf.NewConf().Upyun.Bucket,
		NewFile().Name(),
	)
	err := up.Put(&upyun.PutObjectConfig{
		Path:   path,
		Reader: bytes.NewBuffer(body),
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", host, path), nil
}
