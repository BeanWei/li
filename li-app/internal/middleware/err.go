package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

// ErrorHandler 用来兜底的错误处理
func ErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		r.Response.ClearBuffer()
		var (
			code = gerror.Code(err)
			msg  string
		)
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
			msg = err.Error()
		} else {
			code = gcode.CodeInternalError
			msg = "Oops, 服务器居然开小差了, 请稍后再试吧"
		}
		r.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    code.Code(),
			Message: msg,
		})
	}
}
