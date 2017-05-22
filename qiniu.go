package image

import (
	"bytes"
	"context"
	"fmt"

	"github.com/qiniu/api.v7/kodo"
	"github.com/tiantour/conf"
)

// Qiniu qiniu
type Qiniu struct{}

// NewQiniu new qiniu
func NewQiniu() *Qiniu {
	return &Qiniu{}
}

// init
func init() {
	kodo.SetMac(conf.NewConf().Qiniu.AccessKey, conf.NewConf().Qiniu.SecretKey)
}

// Net net upload
func (q Qiniu) Net(url string) (string, error) {
	body, err := NewFile().Net(url)
	if err != nil {
		return "", err
	}
	return q.Local(body)
}

// Local local upload
func (q Qiniu) Local(body []byte) (string, error) {
	host := conf.NewConf().Qiniu.Host // host
	path := fmt.Sprintf("%s/%s",
		conf.NewConf().Qiniu.Bucket,
		NewFile().Name(),
	)
	zone := 0                                       // 您空间(Bucket)所在的区域
	c := kodo.New(zone, nil)                        // 用默认配置创建 Client
	bucket := c.Bucket(conf.NewConf().Qiniu.Bucket) // 空间
	ctx := context.Background()
	data := bytes.NewBuffer(body)                      // io.reader
	size := int64(len(body))                           // 长度
	err := bucket.Put(ctx, nil, path, data, size, nil) // 上传
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", host, path), nil
}
