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

func (b *textBuilder) Required() *ui.Schema {
	b.schema.Required = true
	return b.schema
}

func (b *textBuilder) MinLength(min int) *ui.Schema {
	b.schema.MinLength = min
	return b.schema
}

func (b *textBuilder) MaxLength(max int) *ui.Schema {
	b.schema.MaxLength = max
	return b.schema
}

func (b *textBuilder) Title(title string) *ui.Schema {
	b.schema.Title = title
	return b.schema
}
