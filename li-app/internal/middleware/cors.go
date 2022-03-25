package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// CORS 允许跨域请求中间件
func CORS(r *ghttp.Request) {
	corsOpts := r.Response.DefaultCORSOptions()
	corsOpts.AllowDomain = g.Cfg().MustGet(r.Context(), "server.corsAllowDomain").Strings()
	r.Response.CORS(corsOpts)
	r.Middleware.Next()
}
