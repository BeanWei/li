package node

import (
	"github.com/BeanWei/li/li-engine/controller"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ActionFormDrawer(name string) *actionformdrawerBuilder {
	return &actionformdrawerBuilder{
		&NodeBuilder{
			schema: &ui.Schema{
				Name:       name,
				Type:       ui.SchemaTypeVoid,
				XComponent: ui.ComponentActionFormDrawer,
				XComponentProps: map[string]interface{}{
					"drawerProps": make(map[string]interface{}),
					"layoutProps": make(map[string]interface{}),
				},
			},
		},
	}
}

type actionformdrawerBuilder struct {
	*NodeBuilder
}

func (b *actionformdrawerBuilder) ForInit(operation string, handler interface{}) *actionformdrawerBuilder {
	b.schema.XComponentProps["forInit"] = operation
	controller.Bind(operation, handler)
	return b
}

func (b *actionformdrawerBuilder) ForSubmit(operation string, handler interface{}) *actionformdrawerBuilder {
	b.schema.XComponentProps["forSubmit"] = operation
	controller.Bind(operation, handler)
	return b
}

func (b *actionformdrawerBuilder) ButtonStyle(style map[string]interface{}) *actionformdrawerBuilder {
	b.schema.XComponentProps["style"] = style
	return b
}

func (b *actionformdrawerBuilder) ButtonType(typ string) *actionformdrawerBuilder {
	b.schema.XComponentProps["type"] = typ
	return b
}

func (b *actionformdrawerBuilder) ButtonStatus(status string) *actionformdrawerBuilder {
	b.schema.XComponentProps["status"] = status
	return b
}

func (b *actionformdrawerBuilder) ButtonSize(size string) *actionformdrawerBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *actionformdrawerBuilder) ButtonShape(shape string) *actionformdrawerBuilder {
	b.schema.XComponentProps["shape"] = shape
	return b
}

func (b *actionformdrawerBuilder) ButtonIcon(icon string) *actionformdrawerBuilder {
	b.schema.XComponentProps["icon"] = icon
	return b
}

func (b *actionformdrawerBuilder) ButtonIconOnly() *actionformdrawerBuilder {
	b.schema.XComponentProps["iconOnly"] = true
	return b
}

func (b *actionformdrawerBuilder) ButtonLong() *actionformdrawerBuilder {
	b.schema.XComponentProps["long"] = true
	return b
}
