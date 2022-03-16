package cmd

import (
	"context"

	"github.com/BeanWei/li/li-app/internal/app/admin"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
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
				admin.Init()
				g.Server().Run()
				return
			},
		},
	)
}
