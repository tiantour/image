package image

import (
	"bytes"
	"fmt"

	"github.com/tiantour/conf"
	"github.com/upyun/go-sdk/upyun"
)

var up = upyun.NewUpYun(
	conf.Data.Upyun.Bucket,
	conf.Data.Upyun.Username,
	conf.Data.Upyun.Passwd,
)

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
	host := conf.Data.Upyun.Host                                          // host
	filePath := fmt.Sprintf("%s/%s", conf.Data.Upyun.Bucket, File.name()) // filePath
	data := bytes.NewBuffer(imageByte)                                      // io.reader
	_, err = up.Put(filePath, data, false, map[string]string{})             // 上传
	if err != nil {
		return
	}
	imagePath = fmt.Sprintf("%s/%s", host, filePath) // image path
	return
}
