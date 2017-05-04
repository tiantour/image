package image

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/tiantour/fetch"
)

// Name
func (i *file) name() string {
	now := time.Now().UnixNano()
	name := strconv.FormatInt(now, 10)
	return fmt.Sprintf("file_%s.jpg", name)
}

// Net
func (i *file) net(imageURL string) ([]byte, error) {
	return fetch.Cmd("get", imageURL)
}

// Local
func (i *file) local(imageURL string) ([]byte, error) {
	imageFile, err := os.Open(imageURL)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(imageFile)
}
