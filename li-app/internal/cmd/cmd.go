package cmd

import (
	"context"

	"github.com/BeanWei/li/li-app/internal/app/admin"
	"github.com/BeanWei/li/li-app/internal/data/ent"
	"github.com/BeanWei/li/li-app/internal/data/ent/migrate"
	_ "github.com/BeanWei/li/li-app/internal/data/hook"
	"github.com/BeanWei/li/li-app/internal/middleware"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gsession"
)

var Li = &gcmd.Command{
	Name:  "li",
	Usage: "li-app cli",
}

func init() {
	Li.AddCommand(
		&gcmd.Command{
			Name:  "run",
			Usage: "run app",
			Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
				s := g.Server()
				s.SetConfigWithMap(g.Map{
					"SessionStorage": gsession.NewStorageRedis(g.Redis()),
				})
				s.Use(
					middleware.CORS,
					middleware.Ctx,
					// TODO: I18N
					// middleware.I18N,
				)
				admin.Init()
				s.Run()
				return
			},
		},
		&gcmd.Command{
			Name:  "migrate",
			Usage: "migrate schemas",
			Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
				return ent.LiClient().Schema.Create(ctx,
					migrate.WithForeignKeys(false),
					migrate.WithDropIndex(true),
					migrate.WithDropColumn(true),
				)
			},
		},
	)
}
