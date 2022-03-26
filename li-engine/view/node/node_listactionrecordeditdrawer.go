package node

import (
	"github.com/BeanWei/li/li-engine/controller"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ListActionRecordEditDrawer(name string) *listactionrecordeditdrawerBuilder {
	return &listactionrecordeditdrawerBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeVoid,
			XComponent: ui.ComponentListActionRecordEditDrawer,
			XComponentProps: map[string]interface{}{
				"drawerProps": make(map[string]interface{}),
				"layoutProps": make(map[string]interface{}),
			},
		},
	}}
}

type listactionrecordeditdrawerBuilder struct {
	*NodeBuilder
}

func (b *listactionrecordeditdrawerBuilder) Title(title string) *listactionrecordeditdrawerBuilder {
	b.schema.Title = title
	return b
}

func (b *listactionrecordeditdrawerBuilder) Description(description string) *listactionrecordeditdrawerBuilder {
	b.schema.Description = description
	return b
}

func (b *listactionrecordeditdrawerBuilder) ForInit(operation string, handler interface{}) *listactionrecordeditdrawerBuilder {
	b.schema.XComponentProps["forInit"] = operation
	controller.Bind(operation, handler)
	return b
}

func (b *listactionrecordeditdrawerBuilder) ForSubmit(operation string, handler interface{}) *listactionrecordeditdrawerBuilder {
	b.schema.XComponentProps["forSubmit"] = operation
	controller.Bind(operation, handler)
	return b
}

func (b *listactionrecordeditdrawerBuilder) ButtonStyle(style map[string]interface{}) *listactionrecordeditdrawerBuilder {
	b.schema.XComponentProps["style"] = style
	return b
}

func (b *listactionrecordeditdrawerBuilder) ButtonType(typ string) *listactionrecordeditdrawerBuilder {
	b.schema.XComponentProps["type"] = typ
	return b
}

func (b *listactionrecordeditdrawerBuilder) ButtonStatus(status string) *listactionrecordeditdrawerBuilder {
	b.schema.XComponentProps["status"] = status
	return b
}

func (b *listactionrecordeditdrawerBuilder) ButtonSize(size string) *listactionrecordeditdrawerBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *listactionrecordeditdrawerBuilder) ButtonShape(shape string) *listactionrecordeditdrawerBuilder {
	b.schema.XComponentProps["shape"] = shape
	return b
}

func (b *listactionrecordeditdrawerBuilder) ButtonIcon(icon string) *listactionrecordeditdrawerBuilder {
	b.schema.XComponentProps["icon"] = icon
	return b
}

func (b *listactionrecordeditdrawerBuilder) ButtonIconOnly() *listactionrecordeditdrawerBuilder {
	b.schema.XComponentProps["iconOnly"] = true
	return b
}

func (b *listactionrecordeditdrawerBuilder) ButtonLong() *listactionrecordeditdrawerBuilder {
	b.schema.XComponentProps["long"] = true
	return b
}
