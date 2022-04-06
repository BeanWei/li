package admin

import (
	"github.com/BeanWei/li/li-app/internal/app/admin/consts"
	"github.com/BeanWei/li/li-app/internal/app/admin/controller"
	"github.com/BeanWei/li/li-app/internal/app/admin/view"
	engine "github.com/BeanWei/li/li-engine"
)

func Init() {
	engine.NewApp(&engine.App{
		Title:     "Li Admin",
		Copyright: "Powered by ❤️璃❤️",
		Entry:     consts.AppEntry,
		Menus: []*engine.AppMenu{
			{
				Title: "仪表盘",
				Icon:  "IconDashboard",
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
				Icon:  "IconSettings",
				Children: []*engine.AppMenu{
					{
						Title: "用户管理",
						Page:  new(view.SystemUser),
					},
				},
			},
			{
				Title: "博客应用",
				Icon:  "IconEdit",
				Children: []*engine.AppMenu{
					{
						Title: "文章管理",
						Page:  new(view.BlogPost),
					},
					{
						Title: "标签管理",
						Page:  new(view.BlogTag),
					},
					{
						Title: "评论管理",
						Page:  new(view.BlogComment),
					},
				},
			},
		},
		NavItems: view.NavItems(),
		Binding: &engine.AppBinding{
			GetCurrentUserController: controller.UserProfile,
			SignForm:                 view.SignFormNode(),
		},
	})
}
