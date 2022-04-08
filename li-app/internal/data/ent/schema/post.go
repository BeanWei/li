package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/BeanWei/li/li-engine/contrib/lient"
	"github.com/BeanWei/li/li-engine/contrib/lient/mixin"
	"github.com/BeanWei/li/li-engine/view/node"
)

type Post struct {
	ent.Schema
}

func (Post) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.XID{},
		mixin.Time{},
	}
}

func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("author_id").
			NotEmpty(),
		field.String("title").
			NotEmpty().
			Annotations(
				lient.Annotation{
					ViewSchema: node.TextArea("title").Title("标题").SetRequired(true).Schema(),
					ColumnProps: &lient.ColumnProps{
						Filterable: true,
						Ellipsis:   true,
					},
				},
			),
		field.String("content").
			NotEmpty().
			Annotations(
				lient.Annotation{
					ViewSchema: node.TextArea("content").Title("内容").SetRequired(true).Schema(),
					ColumnProps: &lient.ColumnProps{
						Filterable: true,
						Width:      300,
						Ellipsis:   true,
					},
				},
			),
		field.Strings("reviewer_ids").
			Optional().
			Annotations(
				lient.Annotation{
					ViewSchema: node.RecordSelect("reviewers").
						Title("审核人").
						Multiple(true).
						FieldNamesTitle("nickname").
						FieldNamesAvatar("avatar").
						FieldNamesDescription("email").
						Schema(),
					ColumnProps: &lient.ColumnProps{
						Filterable: true,
					},
					Edge: lient.XEdge("reviewers", User.Type),
				},
			),
	}
}

func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).
			Ref("posts").
			Field("author_id").
			Unique().
			Required().
			Annotations(
				lient.Annotation{
					ViewSchema: node.RecordSelect("author").
						Title("作者").
						FieldNamesTitle("nickname").
						FieldNamesAvatar("avatar").
						FieldNamesDescription("email").
						Schema(),
					ColumnProps: &lient.ColumnProps{
						Filterable: true,
					},
				},
			),
		edge.To("tags", Tag.Type).
			Annotations(
				lient.Annotation{
					ViewSchema: node.RecordSelect("tags").
						Title("标签").
						Multiple(true).
						FieldNamesTitle("label").
						Schema(),
					ColumnProps: &lient.ColumnProps{
						Filterable: true,
					},
				},
			),
		edge.To("comments", Comment.Type),
	}
}
