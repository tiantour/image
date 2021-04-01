package qiniu

import (
	"fmt"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// Token token
type Token struct{}

// NewToken new token
func NewToken() *Token {
	return &Token{}
}

// Token qiniu access token
func (t *Token) Access(bucket string) string {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200

	mac := qbox.NewMac(AccessKey, SecretKey)
	return putPolicy.UploadToken(mac)
}

// Overwrite qiniu overwrite token
func (t *Token) Overwrite(bucket, key string) string {
	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", bucket, key),
	}
	putPolicy.Expires = 7200

	mac := qbox.NewMac(AccessKey, SecretKey)
	return putPolicy.UploadToken(mac)
}
