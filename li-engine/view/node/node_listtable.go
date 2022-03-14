package node

import (
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ListTable(name string) *listtableBuilder {
	return &listtableBuilder{schema: &ui.Schema{
		Name:            name,
		Type:            ui.SchemaTypeArray,
		XComponent:      ui.ComponentListTable,
		XComponentProps: make(map[string]interface{}),
		Properties:      make(map[string]*ui.Schema),
	}}
}

type listtableBuilder struct {
	schema *ui.Schema
}

func (b *listtableBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *listtableBuilder) Columns(elements ...view.Node) *listtableBuilder {
	for _, element := range elements {
		b.schema.Properties[element.Schema().Name] = element.Schema()
	}
	return b
}
