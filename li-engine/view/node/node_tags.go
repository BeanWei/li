package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Tags(name string) *tagsBuilder {
	return &tagsBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeArray,
			XComponent: ui.ComponentInputTag,
			XDecorator: ui.DecoratorFormItem,
		},
	}}
}

type tagsBuilder struct {
	*NodeBuilder
}

func (b *tagsBuilder) Title(title string) *tagsBuilder {
	b.schema.Title = title
	return b
}

func (b *tagsBuilder) Description(description string) *tagsBuilder {
	b.schema.Description = description
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

func (b *tagsBuilder) AllowClear() *tagsBuilder {
	b.schema.XComponentProps["allowClear"] = true
	return b
}

func (b *tagsBuilder) SaveOnBlur() *tagsBuilder {
	b.schema.XComponentProps["saveOnBlur"] = true
	return b
}

func (b *tagsBuilder) DragToSort() *tagsBuilder {
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
