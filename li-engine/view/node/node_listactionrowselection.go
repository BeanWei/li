package node

import (
	"github.com/BeanWei/li/li-engine/control"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ListActionRowSelection(name string) *listactionrowselectionBuilder {
	return &listactionrowselectionBuilder{schema: &ui.Schema{
		Name:       name,
		Type:       ui.SchemaTypeVoid,
		XComponent: ui.ComponentListActionRowSelection,
		XComponentProps: map[string]interface{}{
			"confirmProps": make(map[string]interface{}),
		},
	}}
}

type listactionrowselectionBuilder struct {
	schema *ui.Schema
}

func (b *listactionrowselectionBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *listactionrowselectionBuilder) Title(title string) *listactionrowselectionBuilder {
	b.schema.Title = title
	return b
}

func (b *listactionrowselectionBuilder) AfterReload() *listactionrowselectionBuilder {
	b.schema.XComponentProps["afterReload"] = true
	return b
}

func (b *listactionrowselectionBuilder) ForSubmit(operation string, controller control.Controller) *listactionrowselectionBuilder {
	b.schema.XComponentProps["forSubmit"] = operation
	control.RegisterController(operation, controller)
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
