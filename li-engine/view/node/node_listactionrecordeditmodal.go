package node

import (
	"github.com/BeanWei/li/li-engine/control"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ListActionRecordEditModal(name string) *listactionrecordeditmodalBuilder {
	return &listactionrecordeditmodalBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeVoid,
			XComponent: ui.ComponentListActionRecordEditModal,
			XComponentProps: map[string]interface{}{
				"modalProps":  make(map[string]interface{}),
				"layoutProps": make(map[string]interface{}),
			},
		},
	}}
}

type listactionrecordeditmodalBuilder struct {
	*NodeBuilder
}

func (b *listactionrecordeditmodalBuilder) ForInit(operation string, controller interface{}) *listactionrecordeditmodalBuilder {
	b.schema.XComponentProps["forInit"] = operation
	control.RegisterController(operation, controller)
	return b
}

func (b *listactionrecordeditmodalBuilder) ForSubmit(operation string, controller interface{}) *listactionrecordeditmodalBuilder {
	b.schema.XComponentProps["forSubmit"] = operation
	control.RegisterController(operation, controller)
	return b
}
