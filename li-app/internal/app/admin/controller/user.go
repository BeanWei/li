package controller

import (
	"context"

	"github.com/BeanWei/li/li-app/internal/app/admin/dto"
	"github.com/BeanWei/li/li-app/internal/data"
	"github.com/BeanWei/li/li-app/internal/data/ent"
	"github.com/BeanWei/li/li-app/internal/data/ent/user"
	"github.com/BeanWei/li/li-app/internal/shared"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

const (
	OperationUserSignIn  = "userSignIn"
	OperationUserProfile = "userProfile"
)

// UserSignIn 用户登录
func UserSignIn(ctx context.Context, req *dto.UserSignInReq) (res *dto.UserSignInRes, err error) {
	usr, err := ent.LiClient().User.
		Query().
		Where(user.EmailEQ(req.Email)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, gerror.NewCode(gcode.CodeNotFound, "账号不存在")
		}
		return nil, gerror.Wrapf(err, "查询用户 %s 错误", req.Email)
	}
	res = &dto.UserSignInRes{
		ID: usr.ID,
	}
	res.Sid, err = shared.Session.SetUser(ctx, usr.ID)
	return
}

// UserProfile 当前用户信息
func UserProfile(ctx context.Context) (res *ent.User, err error) {
	ctxUser := shared.Ctx.Get(ctx).User
	res, err = data.User.GetUser(ctx, ctxUser.ID)
	return
}
