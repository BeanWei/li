package view

import (
	"reflect"

	"github.com/BeanWei/li/li-engine/view/ui"
)

type (
	Schema interface {
		Type() string
		Mixin() []Mixin
		Nodes() []Node
	}

	Mixin interface {
		Nodes() []Node
	}

	Node interface {
		Schema() *ui.Schema
	}
)

func ToPage(schema Schema) (string, map[string]interface{}) {
	var (
		properties = make(map[string]interface{})
		nodes      = make([]Node, 0)
	)
	for _, mixin := range schema.Mixin() {
		nodes = append(nodes, mixin.Nodes()...)
	}
	nodes = append(nodes, schema.Nodes()...)
	for _, node := range nodes {
		properties[node.Schema().Name] = node.Schema()
	}
	return reflect.TypeOf(schema).Elem().Name(), map[string]interface{}{
		"type":       "object",
		"properties": properties,
	}
}
