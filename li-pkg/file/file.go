package file

import (
	"context"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
)

// Storage 存储接口
type Storage interface {
	PutObject(ctx context.Context, input *PutObjectInput) (out *PutObjectOutput, err error)
	DeleteObject(ctx context.Context, input *DeleteObjectInput) (err error)
}

type (
	PutObjectInput struct {
		File        *multipart.FileHeader
		BucketName  string
		FileSize    int64
		ContentType string
	}
	PutObjectOutput struct {
		FileName string
		FileUrl  string
	}
	DeleteObjectInput struct {
		BucketName string
		FileName   string
	}
)

func (i *PutObjectInput) filename() string {
	return strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36)+grand.S(6)) + gfile.Ext(i.File.Filename)
}

func (i *PutObjectInput) output(key string) *PutObjectOutput {
	url := "/" + key
	if i.BucketName != "" {
		url = "/" + i.BucketName + url
	}
	return &PutObjectOutput{
		FileName: key,
		FileUrl:  url,
	}
}
