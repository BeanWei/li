package cmd

import (
	"context"

	"github.com/BeanWei/li/li-app/internal/data/ent"
	"github.com/gogf/gf/v2/os/gcmd"
)

func userCreate(ctx context.Context, parser *gcmd.Parser) (err error) {
	err = ent.LiClient().User.Create().
		SetEmail(parser.GetOpt("email").String()).
		SetPassword(parser.GetOpt("password").String()).
		SetNickname(parser.GetOpt("nickname").String()).
		SetIsAdmin(parser.GetOpt("isAdmin").Bool()).
		Exec(ctx)
	return
}
