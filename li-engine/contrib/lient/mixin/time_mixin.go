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
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			Annotations(lient.Annotation{
				DisableCreate: true,
				DisableUpdate: true,
			}),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(lient.Annotation{
				DisableCreate: true,
				DisableUpdate: true,
			}),
		field.Time("deleted_at").
			Nillable().
			Optional().
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
