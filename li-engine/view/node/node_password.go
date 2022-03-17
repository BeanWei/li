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

func (b *passwordBuilder) Required() *passwordBuilder {
	b.schema.Required = true
	return b
}

func (b *passwordBuilder) MinLength(min int) *passwordBuilder {
	b.schema.MinLength = min
	return b
}

func (b *passwordBuilder) MaxLength(max int) *passwordBuilder {
	b.schema.MaxLength = max
	return b
}

func (b *passwordBuilder) Title(title string) *passwordBuilder {
	b.schema.Title = title
	return b
}
