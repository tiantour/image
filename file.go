package image

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/tiantour/requests"
)

// Name
func (i *file) name() (imageName string) {
	now := time.Now().UnixNano()
	name := strconv.FormatInt(now, 10)
	imageName = fmt.Sprintf("file_%s.jpg", name)
	return
}

// Net
func (i *file) net(imageURL string) (imageByte []byte, err error) {
	_, requestData, requestHeader := requests.Options()
	imageByte, err = requests.Get(imageURL, requestData, requestHeader)
	return
}

// Local
func (i *file) local(imageURL string) (imageByte []byte, err error) {
	imageFile, err := os.Open(imageURL)
	imageByte, err = ioutil.ReadAll(imageFile)
	return
}
