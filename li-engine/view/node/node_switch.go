package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Switch(name string) *switchBuilder {
	return &switchBuilder{schema: &ui.Schema{
		Name:       name,
		Type:       ui.SchemaTypeBool,
		XComponent: ui.ComponentSwitch,
		XDecorator: ui.DecoratorFormItem,
	}}
}

type switchBuilder struct {
	schema *ui.Schema
}

func (b *switchBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *switchBuilder) Title(title string) *ui.Schema {
	b.schema.Title = title
	return b.schema
}
