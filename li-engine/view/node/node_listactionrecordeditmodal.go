package node

import (
	"github.com/BeanWei/li/li-engine/control"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ListActionRecordEditModal(name string) *listactionrecordeditmodalBuilder {
	return &listactionrecordeditmodalBuilder{schema: &ui.Schema{
		Name:       name,
		Type:       ui.SchemaTypeVoid,
		XComponent: ui.ComponentListActionRecordEditModal,
		XComponentProps: map[string]interface{}{
			"modalProps":  make(map[string]interface{}),
			"layoutProps": make(map[string]interface{}),
		},
	}}
}

type listactionrecordeditmodalBuilder struct {
	schema *ui.Schema
}

func (b *listactionrecordeditmodalBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *listactionrecordeditmodalBuilder) Title(title string) *listactionrecordeditmodalBuilder {
	b.schema.Title = title
	return b
}

func (b *listactionrecordeditmodalBuilder) Child(elements ...view.Node) *listactionrecordeditmodalBuilder {
	for _, element := range elements {
		b.schema.Properties[element.Schema().Name] = element.Schema()
	}
	return b
}

func (b *listactionrecordeditmodalBuilder) ForInit(operation string, controller control.Controller) *listactionrecordeditmodalBuilder {
	b.schema.XComponentProps["forInit"] = operation
	control.RegisterController(operation, controller)
	return b
}

func (b *listactionrecordeditmodalBuilder) ForSubmit(operation string, controller control.Controller) *listactionrecordeditmodalBuilder {
	b.schema.XComponentProps["forSubmit"] = operation
	control.RegisterController(operation, controller)
	return b
}
