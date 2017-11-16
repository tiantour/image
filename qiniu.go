package image

import (
	"bytes"
	"context"
	"fmt"

	"github.com/qiniu/api.v7/kodo"
	"github.com/tiantour/conf"
)

var (
	cfq = conf.NewConf().Image["qiniu"]
)

func init() {
	kodo.SetMac(cfq.Uname, cfq.Passwd)
}

// Qiniu qiniu
type Qiniu struct{}

// NewQiniu new qiniu
func NewQiniu() *Qiniu {
	return &Qiniu{}
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
	path := fmt.Sprintf("%s/%s",
		cfq.Bucket,
		NewFile().Name(),
	)
	zone := 0                      // 您空间(Bucket)所在的区域
	c := kodo.New(zone, nil)       // 用默认配置创建 Client
	bucket := c.Bucket(cfq.Bucket) // 空间
	ctx := context.Background()
	data := bytes.NewBuffer(body)                                  // io.reader
	err := bucket.Put(ctx, nil, path, data, int64(len(body)), nil) // 上传
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", cfq.Host, path), nil
}
