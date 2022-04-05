package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func Text(name string) *textBuilder {
	return &textBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeString,
			XComponent:      ui.ComponentInput,
			XDecorator:      ui.DecoratorFormItem,
			XComponentProps: make(map[string]interface{}),
		},
	}}
}

type textBuilder struct {
	*NodeBuilder
}

func (b *textBuilder) AC(f ac.AC) *textBuilder {
	b.schema.AC = f
	return b
}

func (b *textBuilder) Title(title string) *textBuilder {
	b.schema.Title = title
	return b
}

func (b *textBuilder) Description(description string) *textBuilder {
	b.schema.Description = description
	return b
}

func (b *textBuilder) Default(value interface{}) *textBuilder {
	b.schema.Default = value
	return b
}

func (b *textBuilder) AllowClear(allowClear bool) *textBuilder {
	b.schema.XComponentProps["allowClear"] = allowClear
	return b
}

func (b *textBuilder) Placeholder(placeholder string) *textBuilder {
	b.schema.XComponentProps["placeholder"] = placeholder
	return b
}

func (b *textBuilder) AddBefore(addBefore string) *textBuilder {
	b.schema.XComponentProps["addBefore"] = addBefore
	return b
}

func (b *textBuilder) AddAfter(addAfter string) *textBuilder {
	b.schema.XComponentProps["addAfter"] = addAfter
	return b
}

func (b *textBuilder) Prefix(prefix string) *textBuilder {
	b.schema.XComponentProps["prefix"] = prefix
	return b
}

func (b *textBuilder) Suffix(suffix string) *textBuilder {
	b.schema.XComponentProps["suffix"] = suffix
	return b
}

func (b *textBuilder) BeforeStyle(beforeStyle map[string]interface{}) *textBuilder {
	b.schema.XComponentProps["beforeStyle"] = beforeStyle
	return b
}

func (b *textBuilder) AfterStyle(afterStyle map[string]interface{}) *textBuilder {
	b.schema.XComponentProps["afterStyle"] = afterStyle
	return b
}

func (b *textBuilder) Size(size string) *textBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *textBuilder) ShowWordLimit(showWordLimit bool) *textBuilder {
	b.schema.XComponentProps["showWordLimit"] = showWordLimit
	return b
}
