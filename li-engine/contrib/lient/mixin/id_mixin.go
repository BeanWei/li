package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/BeanWei/li/li-engine/contrib/lient"
	"github.com/rs/xid"
)

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
				return xid.New().String()
			}).
			Annotations(lient.Annotation{
				DisableCreate: true,
			}),
	}
}

var _ ent.Mixin = (*XID)(nil)
