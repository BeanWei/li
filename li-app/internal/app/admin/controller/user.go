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
	user, err := ent.LiClient().User.
		Query().
		Where(user.EmailEQ(req.Passport)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, gerror.WrapCode(gcode.CodeNotFound, err, "账号不存在")
		}
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err)
	}
	res.ID = user.ID
	res.Sid, err = shared.Session.SetUser(ctx, user.ID)
	return
}

// UserProfile 当前用户信息
func UserProfile(ctx context.Context, req *dto.UserProfileReq) (res *dto.UserProfileRes, err error) {
	ctxUser := shared.Ctx.Get(ctx).User
	res.User, err = data.User.GetUser(ctx, ctxUser.ID)
	return
}
