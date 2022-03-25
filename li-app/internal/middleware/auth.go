package middleware

import (
	"github.com/BeanWei/li/li-app/internal/app/admin/controller"
	"github.com/BeanWei/li/li-app/internal/shared"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Authentication 登录验证
func Authentication(r *ghttp.Request) {
	operation := r.GetForm("operation").String()
	if operation != controller.OperationUserSignIn {
		ctxUser := shared.Ctx.Get(r.Context()).User
		if ctxUser == nil {
			r.Response.WriteJson(ghttp.DefaultHandlerResponse{
				Code:    gcode.CodeNotAuthorized.Code(),
				Message: "请登录",
			})
		}
	}
	r.Middleware.Next()
}
