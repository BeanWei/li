package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ArrayTableAddition(name string) *arraytableadditionBuilder {
	return &arraytableadditionBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentArrayTableAddition,
			XComponentProps: make(map[string]interface{}),
		},
	}}
}

type arraytableadditionBuilder struct {
	*NodeBuilder
}

func (b *arraytableadditionBuilder) AC(f ac.AC) *arraytableadditionBuilder {
	b.schema.AC = f
	return b
}

func (b *arraytableadditionBuilder) Title(title string) *arraytableadditionBuilder {
	b.schema.XComponentProps["title"] = title
	return b
}

func (b *arraytableadditionBuilder) Method(method string) *arraytableadditionBuilder {
	b.schema.XComponentProps["method"] = method
	return b
}

func (b *arraytableadditionBuilder) DefaultValue(defaultValue map[string]interface{}) *arraytableadditionBuilder {
	b.schema.XComponentProps["defaultValue"] = defaultValue
	return b
}

func (b *arraytableadditionBuilder) ButtonStyle(style map[string]interface{}) *arraytableadditionBuilder {
	b.schema.XComponentProps["style"] = style
	return b
}

func (b *arraytableadditionBuilder) ButtonType(typ string) *arraytableadditionBuilder {
	b.schema.XComponentProps["type"] = typ
	return b
}

func (b *arraytableadditionBuilder) ButtonStatus(status string) *arraytableadditionBuilder {
	b.schema.XComponentProps["status"] = status
	return b
}

func (b *arraytableadditionBuilder) ButtonSize(size string) *arraytableadditionBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *arraytableadditionBuilder) ButtonShape(shape string) *arraytableadditionBuilder {
	b.schema.XComponentProps["shape"] = shape
	return b
}

func (b *arraytableadditionBuilder) ButtonIcon(icon string) *arraytableadditionBuilder {
	b.schema.XComponentProps["icon"] = icon
	return b
}

func (b *arraytableadditionBuilder) ButtonIconOnly(iconOnly bool) *arraytableadditionBuilder {
	b.schema.XComponentProps["iconOnly"] = iconOnly
	return b
}

func (b *arraytableadditionBuilder) ButtonLong(long bool) *arraytableadditionBuilder {
	b.schema.XComponentProps["long"] = long
	return b
}
