package mixin

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/BeanWei/li/li-engine/contrib/lient"
)

type Time struct {
	mixin.Schema
}

func (Time) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("created_at").
			Immutable().
			DefaultFunc(func() int64 {
				return time.Now().Unix()
			}).
			Annotations(lient.Annotation{
				DisableCreate: true,
				DisableUpdate: true,
			}),
		field.Int64("updated_at").
			DefaultFunc(func() int64 {
				return time.Now().Unix()
			}).
			UpdateDefault(func() int64 {
				return time.Now().Unix()
			}).
			Annotations(lient.Annotation{
				DisableCreate: true,
				DisableUpdate: true,
			}),
		field.Int64("deleted_at").
			Optional().
			Default(0).
			StructTag(`json:"-"`).
			Annotations(lient.Annotation{
				DisableCreate: true,
				DisableRead:   true,
				DisableUpdate: true,
			}),
	}
}

func (Time) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
		index.Fields("deleted_at"),
	}
}

var _ ent.Mixin = (*Time)(nil)
