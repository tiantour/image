package image

import (
	"bytes"
	"context"
	"fmt"

	"github.com/qiniu/api.v7/v7/storage"
)

// Qiniu qiniu
type Qiniu struct{}

// NewQiniu new qiniu
func NewQiniu() *Qiniu {
	return &Qiniu{}
}

// Local local upload
func (q *Qiniu) Local(args *File) (string, error) {
	args.Name = NewFormat().Name(args)
	args.Path = NewFormat().Path(args)

	err := q.do(args)
	if err != nil {
		return "", err
	}

	path := fmt.Sprintf("%s/%s", cfq.Host, args.Path)
	return path, nil
}

// do
func (q *Qiniu) do(args *File) error {
	putPolicy := storage.PutPolicy{
		Scope: cfq.Bucket,
	}
	upToken := putPolicy.UploadToken(mac)

	formUploader := storage.NewFormUploader(&storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseHTTPS:      false,
		UseCdnDomains: false,
	})

	return formUploader.Put(
		context.Background(),
		&storage.PutRet{},
		upToken,
		args.Path,
		bytes.NewReader(args.Body),
		args.Size,
		&storage.PutExtra{},
	)
}
