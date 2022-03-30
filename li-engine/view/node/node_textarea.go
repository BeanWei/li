package node

import "github.com/BeanWei/li/li-engine/view/ui"

func TextArea(name string) *textareaBuilder {
	return &textareaBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			XPath:      name,
			Type:       ui.SchemaTypeString,
			XComponent: ui.ComponentInputTextArea,
			XDecorator: ui.DecoratorFormItem,
		},
	}}
}

type textareaBuilder struct {
	*NodeBuilder
}

func (b *textareaBuilder) Title(title string) *textareaBuilder {
	b.schema.Title = title
	return b
}

func (b *textareaBuilder) Description(description string) *textareaBuilder {
	b.schema.Description = description
	return b
}

func (b *textareaBuilder) Default(value interface{}) *textareaBuilder {
	b.schema.Default = value
	return b
}

func (b *textareaBuilder) Style(style map[string]interface{}) *textareaBuilder {
	b.schema.XComponentProps["style"] = style
	return b
}

func (b *textareaBuilder) AllowClear() *textareaBuilder {
	b.schema.XComponentProps["allowClear"] = true
	return b
}

func (b *textareaBuilder) Placeholder(placeholder string) *textareaBuilder {
	b.schema.XComponentProps["placeholder"] = placeholder
	return b
}

func (b *textareaBuilder) Size(size string) *textareaBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *textareaBuilder) ShowWordLimit() *textareaBuilder {
	b.schema.XComponentProps["showWordLimit"] = true
	return b
}

func (b *textareaBuilder) AutoSize() *textareaBuilder {
	b.schema.XComponentProps["autoSize"] = true
	return b
}
