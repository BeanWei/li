package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Checkbox(name string) *checkboxBuilder {
	return &checkboxBuilder{schema: &ui.Schema{
		Name:       name,
		Type:       ui.SchemaTypeBool,
		XComponent: ui.ComponentCheckbox,
	}}
}

type checkboxBuilder struct {
	schema *ui.Schema
}

func (b *checkboxBuilder) Schema() *ui.Schema {
	return b.schema
}
