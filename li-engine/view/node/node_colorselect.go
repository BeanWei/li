package node

import "github.com/BeanWei/li/li-engine/view/ui"

func ColorSelect(name string) *colorselectBuilder {
	return &colorselectBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			XPath:      name,
			Type:       ui.SchemaTypeString,
			XComponent: ui.ComponentColorSelect,
			XDecorator: ui.DecoratorFormItem,
		},
	}}
}

type colorselectBuilder struct {
	*NodeBuilder
}

func (b *colorselectBuilder) Title(title string) *colorselectBuilder {
	b.schema.Title = title
	return b
}

func (b *colorselectBuilder) Description(description string) *colorselectBuilder {
	b.schema.Description = description
	return b
}

func (b *colorselectBuilder) Default(value interface{}) *colorselectBuilder {
	b.schema.Default = value
	return b
}
