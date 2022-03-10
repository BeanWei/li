package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Checkbox(name string) *nodeCheckboxBuilder {
	return &nodeCheckboxBuilder{schema: &ui.Schema{
		Name:       name,
		Type:       ui.SchemaTypeBool,
		XComponent: ui.ComponentCheckbox,
	}}
}

type nodeCheckboxBuilder struct {
	schema *ui.Schema
}

func (b *nodeCheckboxBuilder) Schema() *ui.Schema {
	return b.schema
}
