# image

upload image to server,qiniu,upaiyun with go

qiniu demo

```package main

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
```

upyun demo

```package main

import (
	"fmt"

	"github.com/tiantour/image/upyun"
)

func main() {
	upyun.Operator = "your operator"
	upyun.Password = "your password"
	err := upyun.NewUpload().FromFile(&upyun.Upyun{
		Bucket:    "your bucket",
		Key:       "your file name",
		LocalPath: "your file path",
	})
	fmt.Println(err)
}```

