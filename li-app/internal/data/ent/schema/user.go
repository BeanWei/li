package schema

import (
	"context"

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
		field.String("email").
			NotEmpty().
			Unique().
			Annotations(
				lient.Annotation{
					ViewSchema: node.Email("email").Title("邮箱").Schema(),
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
					ViewSchema:    node.Password("password").Title("密码").Schema(),
					ValidateRule:  "required|password",
					DisableRead:   true,
					DisableUpdate: true,
				},
			),
		field.String("salt").
			Sensitive().
			Comment("密码盐").
			Annotations(
				lient.Annotation{
					DisableCreate: true,
					DisableUpdate: true,
				},
			),
		field.String("nickname").
			NotEmpty().
			Annotations(
				lient.Annotation{
					ViewSchema: node.Text("nickname").Title("昵称").Schema(),
					ColumnProps: &lient.ColumnProps{
						Filterable: true,
					},
				},
			),
		field.String("avatar").
			Optional().
			Annotations(
				lient.Annotation{
					ViewSchema: node.UploadAvatar("avatar").Title("头像").Schema(),
					ColumnProps: &lient.ColumnProps{
						Filterable: false,
						Sortable:   false,
					},
				},
			),
		field.Bool("is_admin").
			Optional().
			Annotations(
				lient.Annotation{
					ViewSchema: node.Checkbox("is_admin").Title("管理员").Schema(),
					ColumnProps: &lient.ColumnProps{
						Filterable: true,
					},
				},
			),
		field.Strings("roles").
			Optional().
			Annotations(
				lient.Annotation{
					ViewSchema: node.Select("roles").
						Title("角色").
						Multiple().
						Option(
							RoleSystemManager, "系统管理员",
						).Schema(),
					ColumnProps: &lient.ColumnProps{
						Filterable: true,
					},
				},
			),
	}
}

func (User) ACL() map[string]ac.AC {
	return map[string]ac.AC{
		"@listUser":       nil,
		"@addUser":        nil,
		"@updateUser":     nil,
		"@deleteManyUser": nil,
		"@deleteUser": func(ctx context.Context) (pass bool, err error) {
			return false, nil
		},
	}
}
