package image

import (
	"bytes"
	"context"
	"fmt"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"github.com/tiantour/conf"
)

var (
	mac *qbox.Mac
	cfq = conf.NewImage().Data["qiniu"]
)

func init() {
	mac = qbox.NewMac(cfq.Uname, cfq.Passwd)
}

// Qiniu qiniu
type Qiniu struct{}

// NewQiniu new qiniu
func NewQiniu() *Qiniu {
	return &Qiniu{}
}

// Net net upload
func (q *Qiniu) Net(url string) (string, error) {
	body, err := NewFile().Net(url)
	if err != nil {
		return "", err
	}
	return q.Local(body)
}

// Local local upload
func (q *Qiniu) Local(body []byte) (string, error) {
	putPolicy := storage.PutPolicy{
		Scope: cfq.Bucket,
	}
	upToken := putPolicy.UploadToken(mac)
	formUploader := storage.NewFormUploader(&storage.Config{
		Zone:          &storage.Zone_z0,
		UseHTTPS:      false,
		UseCdnDomains: false,
	})
	key := NewFile().Name()
	data := bytes.NewReader(body)
	dataLen := int64(len(body))
	err := formUploader.Put(context.Background(), &storage.PutRet{}, upToken, key, data, dataLen, &storage.PutExtra{}) // 上传
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", cfq.Host, key), nil
}
