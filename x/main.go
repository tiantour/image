package main

import (
	"fmt"

	"github.com/tiantour/image/qiniu"
)

func main() {
	qiniu.AccessKey = "your access key"
	qiniu.SecretKey = "your secret key"
	err := qiniu.NewUpload().FromFile(&qiniu.Qiniu{
		Bucket:    "your bucket",
		Key:       "your file name",
		LocalPath: "your file path",
	})
	fmt.Println(err)
}
