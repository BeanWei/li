package node

import (
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ListAction(name string) *listactionBuilder {
	return &listactionBuilder{schema: &ui.Schema{
		Name:       name,
		Type:       ui.SchemaTypeVoid,
		XComponent: ui.ComponentListAction,
		Properties: make(map[string]*ui.Schema),
	}}
}

type listactionBuilder struct {
	schema *ui.Schema
}

func (b *listactionBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *listactionBuilder) Child(elements ...view.Node) *listactionBuilder {
	for _, element := range elements {
		b.schema.Properties[element.Schema().Name] = element.Schema()
	}
	return b
}
