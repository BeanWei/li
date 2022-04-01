package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func Switch(name string) *switchBuilder {
	return &switchBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeBool,
			XComponent:      ui.ComponentSwitch,
			XComponentProps: make(map[string]interface{}),
			XDecorator:      ui.DecoratorFormItem,
		},
	}}
}

type switchBuilder struct {
	*NodeBuilder
}

func (b *switchBuilder) AC(f ac.AC) *switchBuilder {
	b.schema.AC = f
	return b
}

func (b *switchBuilder) Title(title string) *switchBuilder {
	b.schema.Title = title
	return b
}

func (b *switchBuilder) Description(description string) *switchBuilder {
	b.schema.Description = description
	return b
}

func (b *switchBuilder) Default(value interface{}) *switchBuilder {
	b.schema.Default = value
	return b
}

func (b *switchBuilder) Size(size string) *switchBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *switchBuilder) Type(typ string) *switchBuilder {
	b.schema.XComponentProps["type"] = typ
	return b
}

func (b *switchBuilder) CheckedText(checkedText string) *switchBuilder {
	b.schema.XComponentProps["checkedText"] = checkedText
	return b
}

func (b *switchBuilder) UncheckedText(uncheckedText string) *switchBuilder {
	b.schema.XComponentProps["uncheckedText"] = uncheckedText
	return b
}

func (b *switchBuilder) UncheckedIcon(uncheckedIcon string) *switchBuilder {
	b.schema.XComponentProps["UncheckedIcon"] = uncheckedIcon
	return b
}

func (b *switchBuilder) CheckedIcon(checkedIcon string) *switchBuilder {
	b.schema.XComponentProps["checkedIcon"] = checkedIcon
	return b
}
