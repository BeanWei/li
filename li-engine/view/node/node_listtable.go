package node

import (
	"github.com/BeanWei/li/li-engine/view/data"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ListTable(name string) *listtableBuilder {
	return &listtableBuilder{schema: &ui.Schema{
		Name:            name,
		Type:            ui.SchemaTypeVoid,
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

func (b *listtableBuilder) DataProvider(operation string, handler data.Handler) *listtableBuilder {
	b.schema.XOperation = operation
	data.RegisterHandler(operation, handler)
	return b
}
