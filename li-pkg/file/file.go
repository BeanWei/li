package file

import (
	"context"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
)

// Storage 存储接口
type Storage interface {
	PutObject(ctx context.Context, input *PutObjectInput) (out *PutObjectOutput, err error)
	DeleteObject(ctx context.Context, input *DeleteObjectInput) (err error)
	// TODO: 代理层实现 图片裁剪 + office 预览
	Proxy(r *ghttp.Request, input *ProxyInput)
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
	ProxyInput struct {
		BucketName string
		FileName   string
	}
)

func (i *PutObjectInput) filename() string {
	return strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36)+grand.S(6)) + gfile.Ext(i.File.Filename)
}

var (
	ErrorDisabled = gerror.NewOption(gerror.Option{
		Text: "this feature is disabled in this storage",
		Code: gcode.CodeNotSupported,
	})
)
