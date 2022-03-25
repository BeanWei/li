package node

import (
	"github.com/BeanWei/li/li-engine/controller"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ListActionRowSelection(name string) *listactionrowselectionBuilder {
	return &listactionrowselectionBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeVoid,
			XComponent: ui.ComponentListActionRowSelection,
			XComponentProps: map[string]interface{}{
				"confirmProps": make(map[string]interface{}),
			},
		},
	}}
}

type listactionrowselectionBuilder struct {
	*NodeBuilder
}

func (b *listactionrowselectionBuilder) AfterReload() *listactionrowselectionBuilder {
	b.schema.XComponentProps["afterReload"] = true
	return b
}

func (b *listactionrowselectionBuilder) ForSubmit(operation string, handler interface{}) *listactionrowselectionBuilder {
	b.schema.XComponentProps["forSubmit"] = operation
	controller.Bind(operation, handler)
	return b
}

func (b *listactionrowselectionBuilder) ButtonStatus(status string) *listactionrowselectionBuilder {
	b.schema.XComponentProps["status"] = status
	return b
}

func (b *listactionrowselectionBuilder) ButtonType(typ string) *listactionrowselectionBuilder {
	b.schema.XComponentProps["type"] = typ
	return b
}

func (b *listactionrowselectionBuilder) ConfirmTitle(title string) *listactionrowselectionBuilder {
	confirmProps, ok := b.schema.XComponentProps["confirmProps"].(map[string]interface{})
	if ok {
		confirmProps["title"] = title
		b.schema.XComponentProps["confirmProps"] = confirmProps
	}
	return b
}
