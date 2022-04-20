package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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
					ViewSchema: node.Text("nickname").Title("昵称").Schema(),
					ColumnProps: &lient.ColumnProps{
						Filterable: true,
					},
					Queryable: true,
				},
			),
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
					Queryable:     true,
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
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
		edge.To("comments", Comment.Type),
	}
}

func (User) ACL() map[string]ac.AC {
	return map[string]ac.AC{
		"list:User":       nil,
		"create:User":     nil,
		"get:User":        nil,
		"update:User":     nil,
		"delete:User":     nil,
		"deleteMany:User": nil,
		"create:email":    nil,
		"read:email":      nil,
		"update:email":    nil,
	}
}
