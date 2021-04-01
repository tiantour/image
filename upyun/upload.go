package upyun

import (
	"bytes"

	"github.com/upyun/go-sdk/upyun"
)

// Upload upload
type Upload struct{}

// NewUpload new Upload
func NewUpload() *Upload {
	return &Upload{}
}

// Local local upload
func (u *Upload) FromFile(args *Upyun) error {
	up := upyun.NewUpYun(&upyun.UpYunConfig{
		Bucket:   args.Bucket,
		Operator: Operator,
		Password: Password,
	})

	err := up.Put(&upyun.PutObjectConfig{
		Path:      args.Key,
		LocalPath: args.LocalPath,
	})
	return err
}

// do
func (u *Upload) FromStream(args *Upyun) error {
	up := upyun.NewUpYun(&upyun.UpYunConfig{
		Bucket:   args.Bucket,
		Operator: Operator,
		Password: Password,
	})

	return up.Put(&upyun.PutObjectConfig{
		Path:   args.Key,
		Reader: bytes.NewReader(args.Body),
	})
}
