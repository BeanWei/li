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
			HandlerNames: make([]string, 0),
		},
	}}
}

type listactionrowselectionBuilder struct {
	*NodeBuilder
}

func (b *listactionrowselectionBuilder) Title(title string) *listactionrowselectionBuilder {
	b.schema.Title = title
	return b
}

func (b *listactionrowselectionBuilder) Description(description string) *listactionrowselectionBuilder {
	b.schema.Description = description
	return b
}

func (b *listactionrowselectionBuilder) AfterReload() *listactionrowselectionBuilder {
	b.schema.XComponentProps["afterReload"] = true
	return b
}

func (b *listactionrowselectionBuilder) ForSubmit(operation string, handler interface{}) *listactionrowselectionBuilder {
	b.schema.XComponentProps["forSubmit"] = operation
	b.schema.HandlerNames = append(b.schema.HandlerNames, operation)
	controller.Bind(operation, handler)
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

func (b *listactionrowselectionBuilder) ButtonStyle(style map[string]interface{}) *listactionrowselectionBuilder {
	b.schema.XComponentProps["style"] = style
	return b
}

func (b *listactionrowselectionBuilder) ButtonType(typ string) *listactionrowselectionBuilder {
	b.schema.XComponentProps["type"] = typ
	return b
}

func (b *listactionrowselectionBuilder) ButtonStatus(status string) *listactionrowselectionBuilder {
	b.schema.XComponentProps["status"] = status
	return b
}

func (b *listactionrowselectionBuilder) ButtonSize(size string) *listactionrowselectionBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *listactionrowselectionBuilder) ButtonShape(shape string) *listactionrowselectionBuilder {
	b.schema.XComponentProps["shape"] = shape
	return b
}

func (b *listactionrowselectionBuilder) ButtonIcon(icon string) *listactionrowselectionBuilder {
	b.schema.XComponentProps["icon"] = icon
	return b
}

func (b *listactionrowselectionBuilder) ButtonIconOnly() *listactionrowselectionBuilder {
	b.schema.XComponentProps["iconOnly"] = true
	return b
}

func (b *listactionrowselectionBuilder) ButtonLong() *listactionrowselectionBuilder {
	b.schema.XComponentProps["long"] = true
	return b
}

func (b *listactionrowselectionBuilder) ButtonPosition(position string) *listactionrowselectionBuilder {
	b.schema.XComponentProps["position"] = position
	return b
}
