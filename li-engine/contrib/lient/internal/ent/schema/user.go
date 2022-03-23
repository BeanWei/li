package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/BeanWei/li/li-engine/contrib/lient"
	"github.com/BeanWei/li/li-engine/view/ui"
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
					View: &ui.Schema{
						XComponent: ui.ComponentInput,
						Properties: map[string]*ui.Schema{
							"c": {
								Name: "123",
							},
						},
					},
				},
			),
		field.String("email").
			NotEmpty().
			Unique(),
		field.String("password").
			Optional(),
	}
}
