package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/controller"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ListActionRecordDelete(name string) *listactionrecorddeleteBuilder {
	return &listactionrecorddeleteBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeVoid,
			XComponent: ui.ComponentListActionRecordDelete,
			XComponentProps: map[string]interface{}{
				"status":       "danger",
				"confirmProps": make(map[string]interface{}),
			},
			HandlerNames: make([]string, 0),
		},
	}}
}

type listactionrecorddeleteBuilder struct {
	*NodeBuilder
}

func (b *listactionrecorddeleteBuilder) AC(f ac.AC) *listactionrecorddeleteBuilder {
	b.schema.AC = f
	return b
}

func (b *listactionrecorddeleteBuilder) Title(title string) *listactionrecorddeleteBuilder {
	b.SetTitle(title)
	return b
}

func (b *listactionrecorddeleteBuilder) Description(description string) *listactionrecorddeleteBuilder {
	b.SetDescription(description)
	return b
}

func (b *listactionrecorddeleteBuilder) ForSubmit(operation string, handler interface{}) *listactionrecorddeleteBuilder {
	b.schema.XComponentProps["forSubmit"] = operation
	b.schema.HandlerNames = append(b.schema.HandlerNames, operation)
	controller.Bind(operation, handler)
	return b
}

func (b *listactionrecorddeleteBuilder) ConfirmTitle(title string) *listactionrecorddeleteBuilder {
	confirmProps, ok := b.schema.XComponentProps["confirmProps"].(map[string]interface{})
	if ok {
		confirmProps["title"] = title
		b.schema.XComponentProps["confirmProps"] = confirmProps
	}
	return b
}

func (b *listactionrecorddeleteBuilder) ButtonStyle(style map[string]interface{}) *listactionrecorddeleteBuilder {
	b.schema.XComponentProps["style"] = style
	return b
}

func (b *listactionrecorddeleteBuilder) ButtonType(typ string) *listactionrecorddeleteBuilder {
	b.schema.XComponentProps["type"] = typ
	return b
}

func (b *listactionrecorddeleteBuilder) ButtonStatus(status string) *listactionrecorddeleteBuilder {
	b.schema.XComponentProps["status"] = status
	return b
}

func (b *listactionrecorddeleteBuilder) ButtonSize(size string) *listactionrecorddeleteBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *listactionrecorddeleteBuilder) ButtonShape(shape string) *listactionrecorddeleteBuilder {
	b.schema.XComponentProps["shape"] = shape
	return b
}

func (b *listactionrecorddeleteBuilder) ButtonIcon(icon string) *listactionrecorddeleteBuilder {
	b.schema.XComponentProps["icon"] = icon
	return b
}

func (b *listactionrecorddeleteBuilder) ButtonIconOnly(iconOnly bool) *listactionrecorddeleteBuilder {
	b.schema.XComponentProps["iconOnly"] = iconOnly
	return b
}

func (b *listactionrecorddeleteBuilder) ButtonLong(long bool) *listactionrecorddeleteBuilder {
	b.schema.XComponentProps["long"] = long
	return b
}

func (b *listactionrecorddeleteBuilder) ButtonPosition(position string) *listactionrecorddeleteBuilder {
	b.schema.XComponentProps["position"] = position
	return b
}
