package cmd

import (
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"

	_ "github.com/BeanWei/li/li-app/internal/data/hook"
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
			Func:  run,
		},
		&gcmd.Command{
			Name:  "migrate",
			Usage: "migrate schemas",
			Func:  migrate,
		},
		&gcmd.Command{
			Name:  "user-create",
			Usage: "create user",
			Func:  userCreate,
			Arguments: []gcmd.Argument{
				{Name: "email", Short: "u"},
				{Name: "password", Short: "p"},
				{Name: "nickname", Short: "n"},
				{Name: "isAdmin", Short: "a"},
			},
		},
	)
}
