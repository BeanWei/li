package view

import (
	"github.com/BeanWei/li/li-app/internal/data/ent/user"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/com"
	"github.com/BeanWei/li/li-engine/view/node"
)

func NavItems() []view.Node {
	return []view.Node{
		com.LangSwitch("navLangSwitch"),
		com.ThemeSwitch("navThemeSwitch"),
		node.DropdownMenu("navUser").
			TriggerPopupStyle(map[string]interface{}{
				"marginRight": "28px",
			}).
			Droplist(
				node.DropdownMenuItem("updateProfile").
					Children(
						node.ActionFormDrawer("action").
							Title("用户设置").
							IsMenuItem(true).
							InitialValues(map[string]interface{}{
								"nickname": `{{global.currentUser.nickname}}`,
							}).
							Body(
								user.NodeNickname(),
							).
							Footer(
								node.ActionFormDrawerCancel("cancel"),
								node.ActionFormDrawerSubmit("submit"),
							),
					),
				node.DropdownMenuItem("updatePassword").Title("修改密码"),
				node.Divider("divider").
					Style(map[string]interface{}{
						"margin": "4px 0",
					}),
				node.DropdownMenuItem("signOut").Title("退出登录"),
			).
			Children(
				node.Avatar("currentUser").
					Size(32).
					Src(`{{global.currentUser.avatar}}`).
					Style(map[string]interface{}{
						"cursor": "pointer",
					}).
					SetXContent(`{{global.currentUser.nickname?.[0] || "Li"}}`),
			),
	}
}
