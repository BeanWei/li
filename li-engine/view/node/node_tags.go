package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func Tags(name string) *tagsBuilder {
	return &tagsBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeArray,
			XComponent:      ui.ComponentInputTag,
			XComponentProps: make(map[string]interface{}),
			XDecorator:      ui.DecoratorFormItem,
		},
	}}
}

type tagsBuilder struct {
	*NodeBuilder
}

func (b *tagsBuilder) AC(f ac.AC) *tagsBuilder {
	b.schema.AC = f
	return b
}

func (b *tagsBuilder) Title(title string) *tagsBuilder {
	b.SetTitle(title)
	return b
}

func (b *tagsBuilder) Description(description string) *tagsBuilder {
	b.SetDescription(description)
	return b
}

func (b *tagsBuilder) Default(value interface{}) *tagsBuilder {
	b.schema.Default = value
	return b
}

func (b *tagsBuilder) Size(size map[string]interface{}) *tagsBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *tagsBuilder) Placeholder(placeholder string) *tagsBuilder {
	b.schema.XComponentProps["placeholder"] = placeholder
	return b
}

func (b *tagsBuilder) AllowClear(allowClear bool) *tagsBuilder {
	b.schema.XComponentProps["allowClear"] = allowClear
	return b
}

func (b *tagsBuilder) SaveOnBlur(saveOnBlur bool) *tagsBuilder {
	b.schema.XComponentProps["saveOnBlur"] = saveOnBlur
	return b
}

func (b *tagsBuilder) DragToSort(dragToSort bool) *tagsBuilder {
	b.schema.XComponentProps["dragToSort"] = true
	return b
}

func (b *tagsBuilder) Suffix(suffix string) *tagsBuilder {
	b.schema.XComponentProps["suffix"] = suffix
	return b
}

func (b *tagsBuilder) Icon(icon string) *tagsBuilder {
	b.schema.XComponentProps["icon"] = icon
	return b
}
