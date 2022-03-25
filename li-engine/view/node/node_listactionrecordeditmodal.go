package node

import (
	"github.com/BeanWei/li/li-engine/controller"
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

func (b *listactionrecordeditmodalBuilder) ForInit(operation string, handler interface{}) *listactionrecordeditmodalBuilder {
	b.schema.XComponentProps["forInit"] = operation
	controller.Bind(operation, handler)
	return b
}

func (b *listactionrecordeditmodalBuilder) ForSubmit(operation string, handler interface{}) *listactionrecordeditmodalBuilder {
	b.schema.XComponentProps["forSubmit"] = operation
	controller.Bind(operation, handler)
	return b
}
