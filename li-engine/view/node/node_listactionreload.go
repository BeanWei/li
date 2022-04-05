package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ListActionReload(name string) *listactionreloadBuilder {
	return &listactionreloadBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeVoid,
			XComponent: ui.ComponentListActionReload,
		},
	}}
}

type listactionreloadBuilder struct {
	*NodeBuilder
}

func (b *listactionreloadBuilder) AC(f ac.AC) *listactionreloadBuilder {
	b.schema.AC = f
	return b
}

func (b *listactionreloadBuilder) Title(title string) *listactionreloadBuilder {
	b.schema.Title = title
	return b
}

func (b *listactionreloadBuilder) Description(description string) *listactionreloadBuilder {
	b.schema.Description = description
	return b
}

func (b *listactionreloadBuilder) Data(data map[string]interface{}) *listactionreloadBuilder {
	b.schema.XComponentProps["data"] = data
	return b
}

func (b *listactionreloadBuilder) ButtonStyle(style map[string]interface{}) *listactionreloadBuilder {
	b.schema.XComponentProps["style"] = style
	return b
}

func (b *listactionreloadBuilder) ButtonType(typ string) *listactionreloadBuilder {
	b.schema.XComponentProps["type"] = typ
	return b
}

func (b *listactionreloadBuilder) ButtonStatus(status string) *listactionreloadBuilder {
	b.schema.XComponentProps["status"] = status
	return b
}

func (b *listactionreloadBuilder) ButtonSize(size string) *listactionreloadBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *listactionreloadBuilder) ButtonShape(shape string) *listactionreloadBuilder {
	b.schema.XComponentProps["shape"] = shape
	return b
}

func (b *listactionreloadBuilder) ButtonIcon(icon string) *listactionreloadBuilder {
	b.schema.XComponentProps["icon"] = icon
	return b
}

func (b *listactionreloadBuilder) ButtonIconOnly(iconOnly bool) *listactionreloadBuilder {
	b.schema.XComponentProps["iconOnly"] = iconOnly
	return b
}

func (b *listactionreloadBuilder) ButtonLong(long bool) *listactionreloadBuilder {
	b.schema.XComponentProps["long"] = long
	return b
}

func (b *listactionreloadBuilder) ButtonPosition(position string) *listactionreloadBuilder {
	b.schema.XComponentProps["position"] = position
	return b
}
