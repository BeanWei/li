package cmd

import (
	"context"

	"github.com/BeanWei/li/li-app/internal/app/admin"
	"github.com/BeanWei/li/li-app/internal/middleware"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gsession"
)

func run(ctx context.Context, parser *gcmd.Parser) (err error) {
	s := g.Server()
	s.SetSessionStorage(gsession.NewStorageRedis(g.Redis()))
	s.Use(
		middleware.ErrorHandler,
		middleware.CORS,
		middleware.Ctx,
		// TODO: I18N
		// middleware.I18N,
		middleware.Authentication,
	)
	admin.Init()
	s.Run()
	return
}
