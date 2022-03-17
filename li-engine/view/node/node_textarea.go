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

func (b *textareaBuilder) Required() *textareaBuilder {
	b.schema.Required = true
	return b
}

func (b *textareaBuilder) MinLength(min int) *textareaBuilder {
	b.schema.MinLength = min
	return b
}

func (b *textareaBuilder) MaxLength(max int) *textareaBuilder {
	b.schema.MaxLength = max
	return b
}

func (b *textareaBuilder) Title(title string) *textareaBuilder {
	b.schema.Title = title
	return b
}
