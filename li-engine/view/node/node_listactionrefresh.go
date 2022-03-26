package node

import "github.com/BeanWei/li/li-engine/view/ui"

func ListActionRefresh(name string) *listactionrefreshBuilder {
	return &listactionrefreshBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeVoid,
			XComponent: ui.ComponentListActionRefresh,
		},
	}}
}

type listactionrefreshBuilder struct {
	*NodeBuilder
}

func (b *listactionrefreshBuilder) ButtonStyle(style map[string]interface{}) *listactionrefreshBuilder {
	b.schema.XComponentProps["style"] = style
	return b
}

func (b *listactionrefreshBuilder) ButtonType(typ string) *listactionrefreshBuilder {
	b.schema.XComponentProps["type"] = typ
	return b
}

func (b *listactionrefreshBuilder) ButtonStatus(status string) *listactionrefreshBuilder {
	b.schema.XComponentProps["status"] = status
	return b
}

func (b *listactionrefreshBuilder) ButtonSize(size string) *listactionrefreshBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *listactionrefreshBuilder) ButtonShape(shape string) *listactionrefreshBuilder {
	b.schema.XComponentProps["shape"] = shape
	return b
}

func (b *listactionrefreshBuilder) ButtonIcon(icon string) *listactionrefreshBuilder {
	b.schema.XComponentProps["icon"] = icon
	return b
}

func (b *listactionrefreshBuilder) ButtonIconOnly() *listactionrefreshBuilder {
	b.schema.XComponentProps["iconOnly"] = true
	return b
}

func (b *listactionrefreshBuilder) ButtonLong() *listactionrefreshBuilder {
	b.schema.XComponentProps["long"] = true
	return b
}
