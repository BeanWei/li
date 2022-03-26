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

func (b *listactionrecordeditdrawerBuilder) ButtonType(typ string) *listactionrecordeditdrawerBuilder {
	b.schema.XComponentProps["type"] = typ
	return b
}
