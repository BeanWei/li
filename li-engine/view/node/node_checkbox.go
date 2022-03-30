package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Checkbox(name string) *checkboxBuilder {
	return &checkboxBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			XPath:      name,
			Type:       ui.SchemaTypeBool,
			XComponent: ui.ComponentCheckbox,
			XDecorator: ui.DecoratorFormItem,
		},
	}}
}

type checkboxBuilder struct {
	*NodeBuilder
}

func (b *checkboxBuilder) Title(title string) *checkboxBuilder {
	b.schema.Title = title
	return b
}

func (b *checkboxBuilder) Description(description string) *checkboxBuilder {
	b.schema.Description = description
	return b
}

func (b *checkboxBuilder) Default(value interface{}) *checkboxBuilder {
	b.schema.Default = value
	return b
}
