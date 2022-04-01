package shared

import "github.com/gogf/gf/v2/container/garray"

// CtxUser 上下文用户信息
type CtxUser struct {
	ID      string
	IsAdmin bool
	Roles   *garray.StrArray
}
