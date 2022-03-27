package view

import (
	"reflect"

	"github.com/BeanWei/li/li-engine/view/ui"
	"github.com/gogf/gf/v2/container/gmap"
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

func ToPage(schema Interface) (string, string) {
	if schema == nil {
		return "", ""
	}

	var (
		properties = gmap.NewListMap()
		nodes      = make([]Node, 0)
	)
	for _, mixin := range schema.Mixin() {
		nodes = append(nodes, mixin.Nodes()...)
	}
	nodes = append(nodes, schema.Nodes()...)
	if len(nodes) == 0 {
		return "", ""
	}
	for _, node := range nodes {
		properties.Set(node.Schema().Name, node.Schema())
	}
	schemaMap := gmap.NewListMap()
	schemaMap.Set("type", "object")
	schemaMap.Set("properties", properties)
	return gstr.CaseKebab(reflect.TypeOf(schema).Elem().Name()), schemaMap.String()
}
