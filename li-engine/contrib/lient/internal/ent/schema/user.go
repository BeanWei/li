package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/contrib/lient"
	"github.com/BeanWei/li/li-engine/contrib/lient/mixin"
	"github.com/BeanWei/li/li-engine/view/node"
)

type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.XID{},
		mixin.Time{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("nickname").
			NotEmpty().
			Annotations(
				lient.Annotation{
					ViewSchema: node.Text("nickname").SetTitle("昵称").Schema(),
					ColumnProps: &lient.ColumnProps{
						Filterable: true,
					},
				},
			),
		field.String("email").
			NotEmpty().
			Unique().
			Annotations(
				lient.Annotation{
					ViewSchema: node.Email("email").SetTitle("邮箱").Schema(),
					ColumnProps: &lient.ColumnProps{
						Filterable: true,
					},
					ValidateRule:  "required|email",
					DisableUpdate: true,
				},
			),
		field.String("password").
			Sensitive().
			Optional().
			Annotations(
				lient.Annotation{
					ViewSchema:    node.Password("password").SetTitle("密码").Schema(),
					ValidateRule:  "required|password",
					DisableRead:   true,
					DisableUpdate: true,
				},
			),
		field.String("avatar").
			Optional().
			Annotations(
				lient.Annotation{
					ViewSchema: node.UploadAvatar("avatar").SetTitle("头像").Schema(),
					ColumnProps: &lient.ColumnProps{
						Filterable: false,
						Sortable:   false,
					},
				},
			),
	}
}

func (User) ACL() map[string]ac.AC {
	return map[string]ac.AC{
		"@listUser":       nil,
		"@addUser":        nil,
		"@deleteManyUser": nil,
		"@deleteUser":     nil,
	}
}
