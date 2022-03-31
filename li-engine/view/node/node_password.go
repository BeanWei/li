package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func Password(name string) *passwordBuilder {
	return &passwordBuilder{&textBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeString,
			XComponent:      ui.ComponentPassword,
			XComponentProps: make(map[string]interface{}),
			XDecorator:      ui.DecoratorFormItem,
		},
	}}}
}

type passwordBuilder struct {
	*textBuilder
}

func (b *passwordBuilder) AC(f ac.AC) *passwordBuilder {
	b.schema.AC = f
	return b
}

func (b *passwordBuilder) Title(title string) *passwordBuilder {
	b.schema.Title = title
	return b
}

func (b *passwordBuilder) Description(description string) *passwordBuilder {
	b.schema.Description = description
	return b
}

func (b *passwordBuilder) Default(value interface{}) *passwordBuilder {
	b.schema.Default = value
	return b
}

func (b *passwordBuilder) VisibilityToggle() *passwordBuilder {
	b.schema.XComponentProps["visibilityToggle"] = true
	return b
}

func (b *passwordBuilder) DefaultVisibility() *passwordBuilder {
	b.schema.XComponentProps["defaultVisibility"] = true
	return b
}
