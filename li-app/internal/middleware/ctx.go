package middleware

import (
	"github.com/BeanWei/li/li-app/internal/data"
	"github.com/BeanWei/li/li-app/internal/shared"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Ctx 自定义上下文变量
func Ctx(r *ghttp.Request) {
	customCtx := &shared.Context{
		Session: r.Session,
	}
	shared.Ctx.Init(r, customCtx)
	uid := shared.Session.GetUserID(r.Context())
	if uid != "" {
		usr, err := data.User.GetUser(r.Context(), uid)
		if err != nil {
			r.Response.WriteJson(ghttp.DefaultHandlerResponse{
				Code:    gerror.Code(err).Code(),
				Message: err.Error(),
			})
		} else {
			if usr != nil {
				customCtx.User = &shared.CtxUser{
					ID:      usr.ID,
					IsAdmin: usr.IsAdmin,
				}
			}
			r.Middleware.Next()
		}
	} else {
		r.Middleware.Next()
	}
}
