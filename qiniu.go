package image

import (
	"bytes"
	"context"
	"fmt"

	"github.com/qiniu/api.v7/kodo"
	"github.com/tiantour/conf"
)

// init
func init() {
	kodo.SetMac(conf.Data.Qiniu.AccessKey, conf.Data.Qiniu.SecretKey)
}

// Net
func (q *qn) Net(imageURL string) (imagePath string, err error) {
	imageByte, err := File.net(imageURL)
	if err != nil {
		return
	}
	bytes.NewBuffer(imageByte)
	imagePath, err = Qiniu.Local(imageByte)
	return
}

//Local
func (q *qn) Local(imageByte []byte) (imagePath string, err error) {
	host := conf.Data.Qiniu.Host                                          // host
	filePath := fmt.Sprintf("%s/%s", conf.Data.Qiniu.Bucket, File.name()) //filePath
	zone := 0                                                               // 您空间(Bucket)所在的区域
	c := kodo.New(zone, nil)                                                // 用默认配置创建 Client
	bucket := c.Bucket(conf.Data.Qiniu.Bucket)                            // 空间
	ctx := context.Background()
	data := bytes.NewBuffer(imageByte)                    // io.reader
	size := int64(len(imageByte))                         // 长度
	err = bucket.Put(ctx, nil, filePath, data, size, nil) // 上传
	if err != nil {
		return
	}
	imagePath = fmt.Sprintf("%s/%s", host, filePath) // image path
	return
}
