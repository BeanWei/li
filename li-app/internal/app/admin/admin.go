package admin

import (
	"github.com/BeanWei/li/li-app/internal/app/admin/view"
	engine "github.com/BeanWei/li/li-engine"
)

func Init() {
	engine.NewApp(&engine.App{
		Title: "Li Admin",
		Menus: []*engine.AppMenu{
			{
				Name: "仪表盘",
				Children: []*engine.AppMenu{
					{
						Name:   "工作台",
						Page:   new(view.WorkplacePage),
						IsHome: true,
					},
				},
			},
		},
	})
}
