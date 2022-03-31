package view

import (
	"reflect"

	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/controller"
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

func recursionXPath(s *ui.Schema) {
	if s.AC != nil {
		ac.Bind(s.XPath, s.AC)
		for _, hn := range s.HandlerNames {
			controller.UseWithSchemaPath(hn, s.XPath)
		}
	} else if len(s.HandlerNames) > 0 {
		// 如果当前节点上定义了 controller 但是没有定义 AC.
		// 则向上查找离的最近的父节点上的 AC.
		pathItems := gstr.Split(gstr.TrimRight(s.XPath, "."+s.Name), ".")
		pathItemsLen := len(pathItems)
		parentPaths := make([]string, pathItemsLen)
		for i := 0; i < pathItemsLen; i++ {
			// 把路径最短(离的最远)的放在最后面
			parentPaths[pathItemsLen-1-i] = gstr.Join(pathItems[0:i+1], ".")
		}
		for _, path := range parentPaths {
			f := ac.Get(path)
			if f != nil {
				for _, hn := range s.HandlerNames {
					controller.UseWithSchemaPath(hn, path)
				}
				break
			}
		}
	}
	if s.Properties != nil {
		for _, p := range s.Properties.Map() {
			if ss, ok := p.(*ui.Schema); ok {
				ss.XPath = s.XPath + ".properties." + ss.Name
				recursionXPath(ss)
			}
		}
	}
	if s.Items != nil && s.Items.Properties != nil {
		for _, p := range s.Items.Properties.Map() {
			if ss, ok := p.(*ui.Schema); ok {
				ss.XPath = s.XPath + ".items.properties." + ss.Name
				recursionXPath(ss)
			}
		}
	}
}

func ToPage(schema Interface) (string, string) {
	if schema == nil {
		return "", ""
	}
	var (
		rootKey    = gstr.CaseKebab(reflect.TypeOf(schema).Elem().Name())
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
		node.Schema().XPath = rootKey + ".properties." + node.Schema().Name
		recursionXPath(node.Schema())
		properties.Set(node.Schema().Name, node.Schema())
	}
	schemaMap := gmap.NewListMap()
	schemaMap.Set("type", "object")
	schemaMap.Set("properties", properties)
	return rootKey, schemaMap.String()
}
