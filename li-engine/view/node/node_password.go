package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Password(name string) *passwordBuilder {
	return &passwordBuilder{schema: &ui.Schema{
		Name:       name,
		Type:       ui.SchemaTypeString,
		XComponent: ui.ComponentPassword,
		XDecorator: ui.DecoratorFormItem,
	}}
}

type passwordBuilder struct {
	schema *ui.Schema
}

func (b *passwordBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *passwordBuilder) Required() *ui.Schema {
	b.schema.Required = true
	return b.schema
}

func (b *passwordBuilder) MinLength(min int) *ui.Schema {
	b.schema.MinLength = min
	return b.schema
}

func (b *passwordBuilder) MaxLength(max int) *ui.Schema {
	b.schema.MaxLength = max
	return b.schema
}

func (b *passwordBuilder) Title(title string) *ui.Schema {
	b.schema.Title = title
	return b.schema
}
