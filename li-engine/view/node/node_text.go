package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Text(name string) *textBuilder {
	return &textBuilder{schema: &ui.Schema{
		Name:       name,
		Type:       ui.SchemaTypeString,
		XComponent: ui.ComponentInput,
		XDecorator: ui.DecoratorFormItem,
	}}
}

type textBuilder struct {
	schema *ui.Schema
}

func (b *textBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *textBuilder) Required() *textBuilder {
	b.schema.Required = true
	return b
}

func (b *textBuilder) MinLength(min int) *textBuilder {
	b.schema.MinLength = min
	return b
}

func (b *textBuilder) MaxLength(max int) *textBuilder {
	b.schema.MaxLength = max
	return b
}

func (b *textBuilder) Title(title string) *textBuilder {
	b.schema.Title = title
	return b
}
