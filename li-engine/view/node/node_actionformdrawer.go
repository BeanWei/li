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
