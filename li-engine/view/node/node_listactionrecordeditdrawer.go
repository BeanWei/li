package node

import (
	"github.com/BeanWei/li/li-engine/control"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ListActionRecordEditDrawer(name string) *listactionrecordeditdrawerBuilder {
	return &listactionrecordeditdrawerBuilder{schema: &ui.Schema{
		Name:       name,
		Type:       ui.SchemaTypeVoid,
		XComponent: ui.ComponentListActionRecordEditDrawer,
		XComponentProps: map[string]interface{}{
			"drawerProps": make(map[string]interface{}),
			"layoutProps": make(map[string]interface{}),
		},
	}}
}

type listactionrecordeditdrawerBuilder struct {
	schema *ui.Schema
}

func (b *listactionrecordeditdrawerBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *listactionrecordeditdrawerBuilder) Title(title string) *listactionrecordeditdrawerBuilder {
	b.schema.Title = title
	return b
}

func (b *listactionrecordeditdrawerBuilder) Children(elements ...view.Node) *listactionrecordeditdrawerBuilder {
	for _, element := range elements {
		b.schema.Properties[element.Schema().Name] = element.Schema()
	}
	return b
}

func (b *listactionrecordeditdrawerBuilder) ForInit(operation string, controller control.Controller) *listactionrecordeditdrawerBuilder {
	b.schema.XComponentProps["forInit"] = operation
	control.RegisterController(operation, controller)
	return b
}

func (b *listactionrecordeditdrawerBuilder) ForSubmit(operation string, controller control.Controller) *listactionrecordeditdrawerBuilder {
	b.schema.XComponentProps["forSubmit"] = operation
	control.RegisterController(operation, controller)
	return b
}
