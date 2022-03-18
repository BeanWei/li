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
				Title: "仪表盘",
				Children: []*engine.AppMenu{
					{
						Title:  "工作台",
						Page:   new(view.DashboardWorkplace),
						IsHome: true,
					},
				},
			},
			{
				Title: "系统管理",
				Children: []*engine.AppMenu{
					{
						Title: "用户管理",
						Page:  new(view.SystemUser),
					},
				},
			},
		},
		NavItems: view.NavItems(),
	})
}
