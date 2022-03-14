package node

import "github.com/BeanWei/li/li-engine/view/ui"

func ColorSelect(name string) *colorselectBuilder {
	return &colorselectBuilder{schema: &ui.Schema{
		Name:       name,
		Type:       ui.SchemaTypeString,
		XComponent: ui.ComponentColorSelect,
		XDecorator: ui.DecoratorFormItem,
	}}
}

type colorselectBuilder struct {
	schema *ui.Schema
}

func (b *colorselectBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *colorselectBuilder) Required() *ui.Schema {
	b.schema.Required = true
	return b.schema
}

func (b *colorselectBuilder) Title(title string) *ui.Schema {
	b.schema.Title = title
	return b.schema
}
