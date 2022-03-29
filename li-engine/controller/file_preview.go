package controller

import (
	"github.com/BeanWei/li/li-engine/contrib/file"
	"github.com/BeanWei/li/li-engine/service"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

type FilePreviewReq struct {
	BucketName string `json:"bucket_name" p:"bucket_name" v:"required"`
	FileName   string `json:"file_name" p:"file_name" v:"required"`
}

func FilePreviw(r *ghttp.Request) {
	var req *FilePreviewReq
	if err := r.Parse(&req); err != nil {
		code := gerror.Code(err)
		r.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    code.Code(),
			Message: err.Error(),
		})
		return
	}
	client, err := service.NewFileClient(r.Context())
	if err != nil {
		r.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    gcode.CodeInternalError.Code(),
			Message: gcode.CodeInternalError.Message(),
		})
		return
	}
	client.ServeFile(r, &file.ProxyInput{
		BucketName: req.BucketName,
		FileName:   req.FileName,
	})
}
