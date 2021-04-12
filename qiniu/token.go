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

	mac := qbox.NewMac(AccessKey, SecretKey)
	return putPolicy.UploadToken(mac)
}

// Overwrite qiniu overwrite token
func (t *Token) Overwrite(bucket, key string) string {
	scope := fmt.Sprintf("%s:%s", bucket, key)
	putPolicy := storage.PutPolicy{
		Scope: scope,
	}

	mac := qbox.NewMac(AccessKey, SecretKey)
	return putPolicy.UploadToken(mac)
}
