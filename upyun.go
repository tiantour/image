package image

import (
	"bytes"
	"fmt"

	"github.com/tiantour/conf"
	"github.com/upyun/go-sdk/upyun"
)

var up = upyun.NewUpYun(&upyun.UpYunConfig{
	Bucket:   conf.Data.Upyun.Bucket,
	Operator: conf.Data.Upyun.Username,
	Password: conf.Data.Upyun.Passwd,
})

//Net
func (u *upy) Net(imageURL string) (imagePath string, err error) {
	imageByte, err := File.net(imageURL)
	if err != nil {
		return
	}
	imagePath, err = Upyun.Local(imageByte)
	return
}

//Local
func (u *upy) Local(imageByte []byte) (imagePath string, err error) {
	host := conf.Data.Upyun.Host
	filePath := fmt.Sprintf("%s/%s", conf.Data.Upyun.Bucket, File.name()) // filePath
	err = up.Put(&upyun.PutObjectConfig{
		Path:   filePath,
		Reader: bytes.NewBuffer(imageByte),
	})
	if err != nil {
		return
	}
	imagePath = fmt.Sprintf("%s/%s", host, filePath) // image path
	return
}
