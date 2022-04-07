package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ArrayItemsAddition(name string) *arrayitemsadditionBuilder {
	return &arrayitemsadditionBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentArrayItemsAddition,
			XComponentProps: make(map[string]interface{}),
		},
	}}
}

type arrayitemsadditionBuilder struct {
	*NodeBuilder
}

func (b *arrayitemsadditionBuilder) AC(f ac.AC) *arrayitemsadditionBuilder {
	b.schema.AC = f
	return b
}

func (b *arrayitemsadditionBuilder) Title(title string) *arrayitemsadditionBuilder {
	b.schema.XComponentProps["title"] = title
	return b
}

func (b *arrayitemsadditionBuilder) Method(method string) *arrayitemsadditionBuilder {
	b.schema.XComponentProps["method"] = method
	return b
}

func (b *arrayitemsadditionBuilder) DefaultValue(defaultValue map[string]interface{}) *arrayitemsadditionBuilder {
	b.schema.XComponentProps["defaultValue"] = defaultValue
	return b
}

func (b *arrayitemsadditionBuilder) ButtonStyle(style map[string]interface{}) *arrayitemsadditionBuilder {
	b.schema.XComponentProps["style"] = style
	return b
}

func (b *arrayitemsadditionBuilder) ButtonType(typ string) *arrayitemsadditionBuilder {
	b.schema.XComponentProps["type"] = typ
	return b
}

func (b *arrayitemsadditionBuilder) ButtonStatus(status string) *arrayitemsadditionBuilder {
	b.schema.XComponentProps["status"] = status
	return b
}

func (b *arrayitemsadditionBuilder) ButtonSize(size string) *arrayitemsadditionBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *arrayitemsadditionBuilder) ButtonShape(shape string) *arrayitemsadditionBuilder {
	b.schema.XComponentProps["shape"] = shape
	return b
}

func (b *arrayitemsadditionBuilder) ButtonIcon(icon string) *arrayitemsadditionBuilder {
	b.schema.XComponentProps["icon"] = icon
	return b
}

func (b *arrayitemsadditionBuilder) ButtonIconOnly(iconOnly bool) *arrayitemsadditionBuilder {
	b.schema.XComponentProps["iconOnly"] = iconOnly
	return b
}

func (b *arrayitemsadditionBuilder) ButtonLong(long bool) *arrayitemsadditionBuilder {
	b.schema.XComponentProps["long"] = long
	return b
}
