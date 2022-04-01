package middleware

import (
	"github.com/BeanWei/li/li-app/internal/app/admin/controller"
	"github.com/BeanWei/li/li-app/internal/shared"
	engine "github.com/BeanWei/li/li-engine"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Authentication 登录验证
func Authentication(r *ghttp.Request) {
	if r.URL.Path == "/api/liql" {
		operation := r.GetForm("operation").String()
		if operation == controller.OperationUserSignIn ||
			operation == engine.OperationGetAppConfig ||
			operation == engine.OperationGetSignView {
			r.Middleware.Next()
		} else {
			ctxUser := shared.Ctx.Get(r.Context()).User
			if ctxUser == nil {
				r.Response.WriteJson(ghttp.DefaultHandlerResponse{
					Code:    gcode.CodeNotAuthorized.Code(),
					Message: "请登录",
				})
			} else {
				r.Middleware.Next()
			}
		}
	} else {
		r.Middleware.Next()
	}
}
