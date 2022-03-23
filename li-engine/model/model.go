package model

import (
	"github.com/BeanWei/li/li-engine/model/edge"
	"github.com/BeanWei/li/li-engine/model/field"
	"github.com/BeanWei/li/li-engine/model/index"
)

type (
	Interface interface {
		Table() string
		Mixin() []Mixin
		Fields() []Field
		Edges() []Edge
		Indexes() []Index
	}

	Mixin interface {
		Fields() []Field
		Edges() []Edge
		Indexes() []Index
	}

	Field interface {
		Descriptor() *field.Descriptor
	}

	Edge interface {
		Descriptor() *edge.Descriptor
	}

	Index interface {
		Descriptor() *index.Descriptor
	}

	Schema struct {
		Interface
	}
)

func (Schema) Table() string    { return "" }
func (Schema) Mixin() []Mixin   { return nil }
func (Schema) Fields() []Field  { return nil }
func (Schema) Edges() []Edge    { return nil }
func (Schema) Indexes() []Index { return nil }
