package node

import (
	"github.com/BeanWei/li/li-engine/control"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ActionForModal(name string) *actionformmodalBuilder {
	return &actionformmodalBuilder{schema: &ui.Schema{
		Name:       name,
		Type:       ui.SchemaTypeVoid,
		XComponent: ui.ComponentActionFormModal,
		XComponentProps: map[string]interface{}{
			"modalProps":  make(map[string]interface{}),
			"layoutProps": make(map[string]interface{}),
		},
	}}
}

type actionformmodalBuilder struct {
	schema *ui.Schema
}

func (b *actionformmodalBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *actionformmodalBuilder) Title(title string) *actionformmodalBuilder {
	b.schema.Title = title
	return b
}

func (b *actionformmodalBuilder) ForInit(operation string, controller interface{}) *actionformmodalBuilder {
	b.schema.XComponentProps["forInit"] = operation
	control.RegisterController(operation, controller)
	return b
}

func (b *actionformmodalBuilder) ForSubmit(operation string, controller interface{}) *actionformmodalBuilder {
	b.schema.XComponentProps["forSubmit"] = operation
	control.RegisterController(operation, controller)
	return b
}
