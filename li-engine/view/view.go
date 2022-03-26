package view

import (
	"reflect"

	"github.com/BeanWei/li/li-engine/view/ui"
	"github.com/gogf/gf/v2/text/gstr"
)

type (
	Interface interface {
		Mixin() []Mixin
		Nodes() []Node
	}

	Mixin interface {
		Nodes() []Node
	}

	Node interface {
		Schema() *ui.Schema
	}

	Schema struct {
		Interface
	}
)

func (Schema) Mixin() []Mixin { return nil }
func (Schema) Nodes() []Node  { return nil }

func ToPage(schema Interface) (string, map[string]interface{}) {
	if schema == nil {
		return "", nil
	}

	var (
		properties = make(map[string]interface{})
		nodes      = make([]Node, 0)
	)
	for _, mixin := range schema.Mixin() {
		nodes = append(nodes, mixin.Nodes()...)
	}
	nodes = append(nodes, schema.Nodes()...)
	if len(nodes) == 0 {
		return "", nil
	}
	for _, node := range nodes {
		properties[node.Schema().Name] = node.Schema()
	}
	return gstr.CaseKebab(reflect.TypeOf(schema).Elem().Name()), map[string]interface{}{
		"type":       "object",
		"properties": properties,
	}
}
