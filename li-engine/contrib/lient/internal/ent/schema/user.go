package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/BeanWei/li/li-engine/contrib/lient"
	"github.com/BeanWei/li/li-engine/view/node"
)

type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("nickname").
			NotEmpty().
			Annotations(
				lient.Annotation{
					ViewSchema: node.Text("nickname").Schema(),
					ColumnProps: &lient.ColumnProps{
						Title:      "昵称",
						Filterable: true,
					},
				},
			),
		field.String("email").
			NotEmpty().
			Unique().
			Annotations(
				lient.Annotation{
					ViewSchema: node.Email("email").Schema(),
					ColumnProps: &lient.ColumnProps{
						Title:      "邮箱",
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
					ViewSchema:    node.Password("password").Schema(),
					ValidateRule:  "password|required",
					DisableRead:   true,
					DisableUpdate: true,
				},
			),
	}
}
