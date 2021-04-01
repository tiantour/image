package upyun

var (
	// Operator operator
	Operator string
	// Password password
	Password string
)

// Upyun upyun
type Upyun struct {
	Bucket    string // 空间
	Key       string // 路径
	Body      []byte // 内容
	LocalPath string // 路径
}
