package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/BeanWei/li/li-engine/contrib/lient"
)

type ID struct {
	mixin.Schema
}

func (ID) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Unique().
			Immutable().
			DefaultFunc(func() int64 {
				return lient.NewID()
			}).
			Annotations(lient.Annotation{
				DisableCreate: true,
			}),
	}
}

type XID struct {
	mixin.Schema
}

func (XID) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			NotEmpty().
			Unique().
			Immutable().
			DefaultFunc(func() string {
				return lient.NewXid()
			}).
			Annotations(lient.Annotation{
				DisableCreate: true,
			}),
	}
}

var _ ent.Mixin = (*ID)(nil)
var _ ent.Mixin = (*XID)(nil)
