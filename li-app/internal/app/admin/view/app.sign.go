package view

import (
	"github.com/BeanWei/li/li-app/internal/app/admin/consts"
	"github.com/BeanWei/li/li-app/internal/app/admin/controller"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/node"
)

func SignFormNode() view.Node {
	return node.Form("signform").
		Children(
			node.Email("email").
				Placeholder("邮箱").
				SetRequired(true),
			node.Password("password").
				Placeholder("密码").
				SetRequired(true),
			node.Submit("submit").
				ButtonLong(true).
				ButtonType("primary").
				ButtonStyle(map[string]interface{}{
					"width": "100%",
				}).
				ForSubmit(controller.OperationUserSignIn, controller.UserSignIn).
				ForSubmitSuccessTo(consts.AppEntry).
				SetXContent("登录"),
		)
}
