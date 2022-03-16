package model

import (
	"reflect"

	"github.com/BeanWei/li/li-engine/model/field"
	"github.com/BeanWei/li/li-engine/model/index"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/node"
)

type (
	Schema interface {
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

func ToView(schema Schema) view.Node {
	nodes := make([]view.Node, 0)
	for _, mixin := range schema.Mixin() {
		for _, mfield := range mixin.Fields() {
			nodes = append(nodes, mfield.Descriptor().View)
		}
	}
	for _, field := range schema.Fields() {
		nodes = append(nodes, field.Descriptor().View)
	}
	return node.List(reflect.TypeOf(schema).Elem().Name()).
		Children(nodes...)
}
