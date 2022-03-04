package entity

import (
	"github.com/BeanWei/li/li-engine/entity/field"
	"github.com/BeanWei/li/li-engine/entity/index"
)

type (
	Schema interface {
		Type() string
		Mixin() []Mixin
		Fields() []Field
		Indexes() []Index
	}

	Mixin interface {
		Fields() []Field
		Indexes() []Index
	}

	Field interface {
		Descriptor() *field.Descriptor
	}

	Index interface {
		Descriptor() *index.Descriptor
	}
)
