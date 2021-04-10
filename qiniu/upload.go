package qiniu

import (
	"bytes"
	"context"
	"os"

	"github.com/qiniu/go-sdk/v7/storage"
)

// Upload upload
type Upload struct{}

// NewUpload new upload
func NewUpload() *Upload {
	return &Upload{}
}

// FromFile upload image from file
func (u *Upload) FromFile(args *Qiniu) error {
	formUploader := storage.NewFormUploader(&storage.Config{
		Zone:          &storage.ZoneHuanan, // 空间对应的机房
		UseHTTPS:      true,                // 是否使用https域名
		UseCdnDomains: true,                // 上传是否使用CDN上传加速
	})
	putExtra := storage.PutExtra{}
	upToken := NewToken().Access(args.Bucket)

	ret := storage.PutRet{}
	return formUploader.PutFile(
		context.Background(), // 请求上下文
		&ret,                 // 上传成功后返回的数据
		upToken,              // 业务服务器颁发的上传凭证
		args.Key,             // 要上传的文件访问路径
		args.LocalPath,       // 是要上传的文件的本地路径
		&putExtra,            // 上传的一些可选项
	)
}

// FromStream upload image from stream
func (u *Upload) FromStream(args *Qiniu) error {
	formUploader := storage.NewFormUploader(&storage.Config{
		Zone:          &storage.ZoneHuanan, // 空间对应的机房
		UseHTTPS:      true,                // 是否使用https域名
		UseCdnDomains: false,               // 上传是否使用CDN上传加速
	})
	putExtra := storage.PutExtra{}
	upToken := NewToken().Access(args.Bucket)

	ret := storage.PutRet{}
	return formUploader.Put(
		context.Background(),       // 请求上下文
		&ret,                       // 上传成功后返回的数据
		upToken,                    // 业务服务器颁发的上传凭证
		args.Key,                   // 要上传的文件访问路径
		bytes.NewReader(args.Body), // 文件内容的访问接口（io.Reader）
		int64(len(args.Body)),      // 要上传的文件大小
		&putExtra,                  // 上传的一些可选项
	)
}

// FromBlock upload image from block
func (u *Upload) FromBlock(args *Qiniu) error {
	resumeUploader := storage.NewResumeUploaderV2(&storage.Config{
		Zone:          &storage.ZoneHuanan, // 空间对应的机房
		UseHTTPS:      true,                // 是否使用https域名
		UseCdnDomains: false,               // 上传是否使用CDN上传加速
	})
	putExtra := storage.RputV2Extra{}
	upToken := NewToken().Access(args.Bucket)

	ret := storage.PutRet{}
	return resumeUploader.PutFile(
		context.Background(), // 请求上下文
		&ret,                 // 上传成功后返回的数据
		upToken,              // 业务服务器颁发的上传凭证
		args.Key,             // 要上传的文件访问路径
		args.LocalPath,       // 是要上传的文件的本地路径
		&putExtra,            // 上传的一些可选项
	)
}

// FromRecord upload image from reacord
func (u *Upload) FromRecord(args *Qiniu) error {
	resumeUploader := storage.NewResumeUploaderV2(&storage.Config{
		Zone:          &storage.ZoneHuanan, // 空间对应的机房
		UseHTTPS:      true,                // 是否使用https域名
		UseCdnDomains: false,               // 上传是否使用CDN上传加速
	})
	recorder, err := storage.NewFileRecorder(os.TempDir())
	if err != nil {
		return err
	}
	putExtra := storage.RputV2Extra{
		Recorder: recorder,
	}
	upToken := NewToken().Access(args.Bucket)

	ret := storage.PutRet{}
	return resumeUploader.PutFile(
		context.Background(), // 请求上下文
		&ret,                 // 上传成功后返回的数据
		upToken,              // 业务服务器颁发的上传凭证
		args.Key,             // 要上传的文件访问路径
		args.LocalPath,       // 是要上传的文件的本地路径
		&putExtra,            // 上传的一些可选项
	)
}
