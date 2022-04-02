package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/controller"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ListActionRecordFormDrawer(name string) *listactionrecordformrawerBuilder {
	return &listactionrecordformrawerBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeVoid,
			XComponent: ui.ComponentListActionRecordFormDrawer,
			XComponentProps: map[string]interface{}{
				"drawerProps": make(map[string]interface{}),
				"layoutProps": make(map[string]interface{}),
			},
			HandlerNames: make([]string, 0),
		},
	}}
}

type listactionrecordformrawerBuilder struct {
	*NodeBuilder
}

func (b *listactionrecordformrawerBuilder) AC(f ac.AC) *listactionrecordformrawerBuilder {
	b.schema.AC = f
	return b
}

func (b *listactionrecordformrawerBuilder) Title(title string) *listactionrecordformrawerBuilder {
	b.schema.Title = title
	return b
}

func (b *listactionrecordformrawerBuilder) Description(description string) *listactionrecordformrawerBuilder {
	b.schema.Description = description
	return b
}

func (b *listactionrecordformrawerBuilder) ForInit(operation string, handler interface{}) *listactionrecordformrawerBuilder {
	b.schema.XComponentProps["forInit"] = operation
	b.schema.HandlerNames = append(b.schema.HandlerNames, operation)
	controller.Bind(operation, handler)
	return b
}

func (b *listactionrecordformrawerBuilder) ButtonStyle(style map[string]interface{}) *listactionrecordformrawerBuilder {
	b.schema.XComponentProps["style"] = style
	return b
}

func (b *listactionrecordformrawerBuilder) ButtonType(typ string) *listactionrecordformrawerBuilder {
	b.schema.XComponentProps["type"] = typ
	return b
}

func (b *listactionrecordformrawerBuilder) ButtonStatus(status string) *listactionrecordformrawerBuilder {
	b.schema.XComponentProps["status"] = status
	return b
}

func (b *listactionrecordformrawerBuilder) ButtonSize(size string) *listactionrecordformrawerBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *listactionrecordformrawerBuilder) ButtonShape(shape string) *listactionrecordformrawerBuilder {
	b.schema.XComponentProps["shape"] = shape
	return b
}

func (b *listactionrecordformrawerBuilder) ButtonIcon(icon string) *listactionrecordformrawerBuilder {
	b.schema.XComponentProps["icon"] = icon
	return b
}

func (b *listactionrecordformrawerBuilder) ButtonIconOnly() *listactionrecordformrawerBuilder {
	b.schema.XComponentProps["iconOnly"] = true
	return b
}

func (b *listactionrecordformrawerBuilder) ButtonLong() *listactionrecordformrawerBuilder {
	b.schema.XComponentProps["long"] = true
	return b
}

func (b *listactionrecordformrawerBuilder) ButtonPosition(position string) *listactionrecordformrawerBuilder {
	b.schema.XComponentProps["position"] = position
	return b
}

func (b *listactionrecordformrawerBuilder) DrawerTitle(title string) *listactionrecordformrawerBuilder {
	drawerProps, ok := b.schema.XComponentProps["drawerProps"].(map[string]interface{})
	if ok {
		drawerProps["title"] = title
		b.schema.XComponentProps["drawerProps"] = drawerProps
	}
	return b
}

func (b *listactionrecordformrawerBuilder) Body(elements ...view.Node) *listactionrecordformrawerBuilder {
	b.Items(elements...)
	return b
}

func (b *listactionrecordformrawerBuilder) Footer(elements ...view.Node) *listactionrecordformrawerBuilder {
	b.Children(elements...)
	return b
}
