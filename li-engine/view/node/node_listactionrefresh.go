package node

import "github.com/BeanWei/li/li-engine/view/ui"

func ListActionRefresh(name string) *listactionrefreshBuilder {
	return &listactionrefreshBuilder{schema: &ui.Schema{
		Name:       name,
		Type:       ui.SchemaTypeVoid,
		XComponent: ui.ComponentListActionRefresh,
	}}
}

type listactionrefreshBuilder struct {
	schema *ui.Schema
}

func (b *listactionrefreshBuilder) Schema() *ui.Schema {
	return b.schema
}
