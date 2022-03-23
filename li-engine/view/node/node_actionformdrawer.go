package node

import (
	"github.com/BeanWei/li/li-engine/control"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ActionFormDrawer(name string) *actionformdrawerBuilder {
	return &actionformdrawerBuilder{schema: &ui.Schema{
		Name:       name,
		Type:       ui.SchemaTypeVoid,
		XComponent: ui.ComponentActionFormDrawer,
		XComponentProps: map[string]interface{}{
			"drawerProps": make(map[string]interface{}),
			"layoutProps": make(map[string]interface{}),
		},
	}}
}

type actionformdrawerBuilder struct {
	schema *ui.Schema
}

func (b *actionformdrawerBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *actionformdrawerBuilder) Title(title string) *actionformdrawerBuilder {
	b.schema.Title = title
	return b
}

func (b *actionformdrawerBuilder) ForInit(operation string, controller interface{}) *actionformdrawerBuilder {
	b.schema.XComponentProps["forInit"] = operation
	control.RegisterController(operation, controller)
	return b
}

func (b *actionformdrawerBuilder) ForSubmit(operation string, controller interface{}) *actionformdrawerBuilder {
	b.schema.XComponentProps["forSubmit"] = operation
	control.RegisterController(operation, controller)
	return b
}
