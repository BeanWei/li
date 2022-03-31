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
			HandlerNames: make([]string, 0),
		},
	}}
}

type listactionrecordeditmodalBuilder struct {
	*NodeBuilder
}

func (b *listactionrecordeditmodalBuilder) Title(title string) *listactionrecordeditmodalBuilder {
	b.schema.Title = title
	return b
}

func (b *listactionrecordeditmodalBuilder) Description(description string) *listactionrecordeditmodalBuilder {
	b.schema.Description = description
	return b
}

func (b *listactionrecordeditmodalBuilder) ForInit(operation string, handler interface{}) *listactionrecordeditmodalBuilder {
	b.schema.XComponentProps["forInit"] = operation
	b.schema.HandlerNames = append(b.schema.HandlerNames, operation)
	controller.Bind(operation, handler)
	return b
}

func (b *listactionrecordeditmodalBuilder) ForSubmit(operation string, handler interface{}) *listactionrecordeditmodalBuilder {
	b.schema.XComponentProps["forSubmit"] = operation
	b.schema.HandlerNames = append(b.schema.HandlerNames, operation)
	controller.Bind(operation, handler)
	return b
}

func (b *listactionrecordeditmodalBuilder) ButtonStyle(style map[string]interface{}) *listactionrecordeditmodalBuilder {
	b.schema.XComponentProps["style"] = style
	return b
}

func (b *listactionrecordeditmodalBuilder) ButtonType(typ string) *listactionrecordeditmodalBuilder {
	b.schema.XComponentProps["type"] = typ
	return b
}

func (b *listactionrecordeditmodalBuilder) ButtonStatus(status string) *listactionrecordeditmodalBuilder {
	b.schema.XComponentProps["status"] = status
	return b
}

func (b *listactionrecordeditmodalBuilder) ButtonSize(size string) *listactionrecordeditmodalBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *listactionrecordeditmodalBuilder) ButtonShape(shape string) *listactionrecordeditmodalBuilder {
	b.schema.XComponentProps["shape"] = shape
	return b
}

func (b *listactionrecordeditmodalBuilder) ButtonIcon(icon string) *listactionrecordeditmodalBuilder {
	b.schema.XComponentProps["icon"] = icon
	return b
}

func (b *listactionrecordeditmodalBuilder) ButtonIconOnly() *listactionrecordeditmodalBuilder {
	b.schema.XComponentProps["iconOnly"] = true
	return b
}

func (b *listactionrecordeditmodalBuilder) ButtonLong() *listactionrecordeditmodalBuilder {
	b.schema.XComponentProps["long"] = true
	return b
}

func (b *listactionrecordeditmodalBuilder) ButtonPosition(position string) *listactionrecordeditmodalBuilder {
	b.schema.XComponentProps["position"] = position
	return b
}
