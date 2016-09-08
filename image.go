package image

// Server Server
var (
	File   = &file{}
	Server = &server{}
	Qiniu  = &qn{}
	Upyun  = &upy{}
)

// Upload Upload
type (
	Upload interface {
		Local(imageByte []byte) (imagePath string, err error)
		Net(imageURL string) (imagePath string, err error)
	}
	file   struct{}
	server struct{}
	qn     struct{}
	upy    struct{}
)
