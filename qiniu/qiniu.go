package qiniu

var (
	// AccessKey Access Key
	AccessKey string

	// SecretKey Secret Key
	SecretKey string
)

// Qiniu Qiniu
type Qiniu struct {
	Bucket    string // 空间
	Key       string // 路径
	Body      []byte // 内容，from stream使用
	LocalPath string // 地址，from file使用
}
