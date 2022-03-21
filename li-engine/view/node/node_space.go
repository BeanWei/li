package node

import (
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func Space(name string) *spaceBuilder {
	return &spaceBuilder{schema: &ui.Schema{
		Name:       name,
		Type:       ui.SchemaTypeVoid,
		XComponent: ui.ComponentSpace,
		Properties: make(map[string]*ui.Schema),
	}}
}

type spaceBuilder struct {
	schema *ui.Schema
}

func (b *spaceBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *spaceBuilder) Children(elements ...view.Node) *spaceBuilder {
	for _, element := range elements {
		b.schema.Properties[element.Schema().Name] = element.Schema()
	}
	return b
}
