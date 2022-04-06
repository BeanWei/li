package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/BeanWei/li/li-engine/contrib/lient"
	"github.com/BeanWei/li/li-engine/contrib/lient/mixin"
	"github.com/BeanWei/li/li-engine/view/node"
)

type Comment struct {
	ent.Schema
}

func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.XID{},
		mixin.Time{},
	}
}

func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("owner_id").
			NotEmpty(),
		field.String("post_id").
			NotEmpty(),
		field.String("content").
			NotEmpty().
			Annotations(
				lient.Annotation{
					ViewSchema: node.TextArea("content").Title("评论内容").SetRequired(true).Schema(),
					ColumnProps: &lient.ColumnProps{
						Filterable: true,
					},
				},
			),
	}
}

func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("comments").
			Field("owner_id").
			Required().
			Unique().
			Annotations(
				lient.Annotation{
					ViewSchema: node.RecordPicker("owner").
						Title("评论者").
						FieldNamesLabel("nickname").
						Schema(),
					ColumnProps: &lient.ColumnProps{
						Filterable: true,
					},
				},
			),
		edge.From("post", Post.Type).
			Ref("comments").
			Field("post_id").
			Required().
			Unique().
			Annotations(
				lient.Annotation{
					ViewSchema: node.RecordPicker("post").
						Title("评论文章").
						FieldNamesLabel("title").
						Schema(),
					ColumnProps: &lient.ColumnProps{
						Filterable: true,
					},
				},
			),
	}
}
