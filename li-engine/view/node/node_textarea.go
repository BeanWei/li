package node

import "github.com/BeanWei/li/li-engine/view/ui"

func TextArea(name string) *textareaBuilder {
	return &textareaBuilder{schema: &ui.Schema{
		Name:       name,
		Type:       ui.SchemaTypeString,
		XComponent: ui.ComponentInputTextArea,
		XDecorator: ui.DecoratorFormItem,
	}}
}

type textareaBuilder struct {
	schema *ui.Schema
}

func (b *textareaBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *textareaBuilder) Required() *ui.Schema {
	b.schema.Required = true
	return b.schema
}

func (b *textareaBuilder) MinLength(min int) *ui.Schema {
	b.schema.MinLength = min
	return b.schema
}

func (b *textareaBuilder) MaxLength(max int) *ui.Schema {
	b.schema.MaxLength = max
	return b.schema
}

func (b *textareaBuilder) Title(title string) *ui.Schema {
	b.schema.Title = title
	return b.schema
}
