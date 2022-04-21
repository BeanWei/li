package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/controller"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ActionFormModalSubmit(name string) *actionformmodalsubmitBuilder {
	return &actionformmodalsubmitBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentActionFormModalSubmit,
			XComponentProps: make(map[string]interface{}),
			HandlerNames:    make([]string, 0),
		},
	}}
}

type actionformmodalsubmitBuilder struct {
	*NodeBuilder
}

func (b *actionformmodalsubmitBuilder) AC(f ac.AC) *actionformmodalsubmitBuilder {
	b.schema.AC = f
	return b
}

func (b *actionformmodalsubmitBuilder) Title(title string) *actionformmodalsubmitBuilder {
	b.SetTitle(title)
	return b
}

func (b *actionformmodalsubmitBuilder) ForSubmit(operation string, handler interface{}) *actionformmodalsubmitBuilder {
	b.schema.XComponentProps["forSubmit"] = operation
	b.schema.HandlerNames = append(b.schema.HandlerNames, operation)
	controller.Bind(operation, handler)
	return b
}
