package image

import (
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/tiantour/conf"
	"github.com/upyun/go-sdk/upyun"
)

// File file
type File struct {
	Body   []byte // 内容
	Name   string // 名称
	Path   string // 路径
	Prefix int64  // 前缀
	Size   int64  // 大小
}

var (
	mac *qbox.Mac
	up  *upyun.UpYun

	cfq = conf.NewImage().Data["qiniu"]
	cfs = conf.NewImage().Data["server"]
	cfu = conf.NewImage().Data["upyun"]
)

func init() {
	mac = qbox.NewMac(cfq.Uname, cfq.Passwd)
	up = upyun.NewUpYun(&upyun.UpYunConfig{
		Bucket:   cfu.Bucket,
		Operator: cfu.Uname,
		Password: cfu.Passwd,
	})

}
