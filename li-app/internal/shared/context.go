package shared

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

// 上下文管理服务
var Ctx = new(ctxShared)

type ctxShared struct{}

// CtxSrvKey 上下文变量存储键名
const ContextKey = "CtxSvcKey"

// Context 请求上下文结构
type Context struct {
	Session *ghttp.Session // session
	User    *CtxUser       // 上下文用户信息
}

// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *ctxShared) Init(r *ghttp.Request, customCtx *Context) {
	r.SetCtxVar(ContextKey, customCtx)
}

// Get 获得上下文变量
func (s *ctxShared) Get(ctx context.Context) *Context {
	value := ctx.Value(ContextKey)
	if value == nil {
		return &Context{}
	}
	if localCtx, ok := value.(*Context); ok {
		return localCtx
	}
	return &Context{}
}

// SetUser 设定上下文用户信息
func (s *ctxShared) SetUser(ctx context.Context, ctxUser *CtxUser) {
	s.Get(ctx).User = ctxUser
}
