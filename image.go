package image

// Upload upload
type Upload interface {
	Local(body []byte) (string, error)
	Net(url string) (string, error)
}
