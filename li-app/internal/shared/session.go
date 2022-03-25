package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
)

// 会话管理服务
var Session = new(sessionShared)

type sessionShared struct{}

const (
	// SessionKeyUser 用户信息session存储键名
	SessionKeyUser = "SessionUserSvcKey"
)

// SetUser 设置用户Session
func (s *sessionShared) SetUser(ctx context.Context, uid string) (sid string, err error) {
	var ttl time.Duration
	ttl, err = gtime.ParseDuration(g.Cfg().MustGet(ctx, "server.sessionCookieMaxAge", "30d").String())
	if err != nil {
		return
	}
	content, err := json.Marshal(map[string]string{
		SessionKeyUser: uid,
	})
	if err != nil {
		return
	}
	r := ghttp.RequestFromCtx(ctx)
	if r == nil {
		return "", gerror.NewCode(gcode.CodeInternalError)
	}
	sid = guid.S([]byte(r.RemoteAddr), []byte(fmt.Sprintf("%v", r.Header)))
	_, err = g.Redis().Do(ctx, "SETEX", sid, int64(ttl.Seconds()), content)
	if err != nil {
		return
	}
	r.Cookie.SetSessionId(sid)
	return
}

// GetUser 获取当前登录的用户ID
func (s *sessionShared) GetUserID(ctx context.Context) string {
	return Ctx.Get(ctx).Session.MustGet(SessionKeyUser).String()
}

// RemoveUser 删除用户Session
func (s *sessionShared) RemoveUser(ctx context.Context) error {
	customCtx := Ctx.Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(SessionKeyUser)
	}
	return nil
}
