package controller

import (
	"context"

	"github.com/BeanWei/li/li-engine/contrib/lifile"
	"github.com/BeanWei/li/li-engine/service"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

type FileUploadRes struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// FileUpload 单文件上传
func FileUpload(ctx context.Context) (res *FileUploadRes, err error) {
	client, err := service.NewFileClient(ctx)
	if err != nil {
		return nil, err
	}

	f := ghttp.RequestFromCtx(ctx).GetUploadFile("file")
	if f == nil {
		err = gerror.NewCode(
			gcode.CodeMissingParameter,
			"file is empty, maybe you retrieve it from invalid field name or form enctype",
		)
		return
	}

	output, err := client.PutObject(ctx, &lifile.PutObjectInput{
		File:       f.FileHeader,
		BucketName: gtime.Now().Format("Ym"),
	})
	if err != nil {
		return nil, err
	}
	res = &FileUploadRes{
		Name: output.FileName,
		URL:  "/upload" + output.FileUrl,
	}
	return
}
