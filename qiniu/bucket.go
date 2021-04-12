package qiniu

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// Bucket bucket
type Bucket struct{}

// NewBucket new bucket
func NewBucket() *Bucket {
	return &Bucket{}
}

// Manager bucket manager
func (b *Bucket) Manager() *storage.BucketManager {
	cfg := storage.Config{
		UseHTTPS: true,
		Zone:     &storage.ZoneHuanan,
	}

	mac := qbox.NewMac(AccessKey, SecretKey)
	return storage.NewBucketManager(mac, &cfg)
}
