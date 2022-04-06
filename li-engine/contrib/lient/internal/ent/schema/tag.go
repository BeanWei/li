package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/BeanWei/li/li-engine/contrib/lient"
	"github.com/BeanWei/li/li-engine/contrib/lient/mixin"
	"github.com/BeanWei/li/li-engine/view/node"
)

type Tag struct {
	ent.Schema
}

func (Tag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.XID{},
		mixin.Time{},
	}
}

func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("label").
			NotEmpty().
			Annotations(
				lient.Annotation{
					ViewSchema: node.Text("label").Title("标签").SetRequired(true).Schema(),
					ColumnProps: &lient.ColumnProps{
						Filterable: true,
					},
				},
			),
	}
}

func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("posts", Post.Type).
			Ref("tags"),
	}
}
