package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/controller"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ActionFormDrawerSubmit(name string) *actionformdrawersubmitBuilder {
	return &actionformdrawersubmitBuilder{
		&NodeBuilder{
			schema: &ui.Schema{
				Name:            name,
				Type:            ui.SchemaTypeVoid,
				XComponent:      ui.ComponentActionFormDrawerSubmit,
				XComponentProps: make(map[string]interface{}),
				HandlerNames:    make([]string, 0),
			},
		},
	}
}

type actionformdrawersubmitBuilder struct {
	*NodeBuilder
}

func (b *actionformdrawersubmitBuilder) AC(f ac.AC) *actionformdrawersubmitBuilder {
	b.schema.AC = f
	return b
}

func (b *actionformdrawersubmitBuilder) Title(title string) *actionformdrawersubmitBuilder {
	b.SetTitle(title)
	return b
}

func (b *actionformdrawersubmitBuilder) ForSubmit(operation string, handler interface{}) *actionformdrawersubmitBuilder {
	b.schema.XComponentProps["forSubmit"] = operation
	b.schema.HandlerNames = append(b.schema.HandlerNames, operation)
	controller.Bind(operation, handler)
	return b
}
