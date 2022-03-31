package node

import (
	"github.com/BeanWei/li/li-engine/controller"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ActionForModal(name string) *actionformmodalBuilder {
	return &actionformmodalBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeVoid,
			XComponent: ui.ComponentActionFormModal,
			XComponentProps: map[string]interface{}{
				"modalProps":  make(map[string]interface{}),
				"layoutProps": make(map[string]interface{}),
			},
			HandlerNames: make([]string, 0),
		},
	}}
}

type actionformmodalBuilder struct {
	*NodeBuilder
}

func (b *actionformmodalBuilder) Title(title string) *actionformmodalBuilder {
	b.schema.Title = title
	return b
}

func (b *actionformmodalBuilder) Description(description string) *actionformmodalBuilder {
	b.schema.Description = description
	return b
}

// TODO: 支持通过回调函数获取
func (b *actionformmodalBuilder) InitialValues(initialValues map[string]interface{}) *actionformmodalBuilder {
	b.schema.XComponentProps["initialValues"] = initialValues
	return b
}

func (b *actionformmodalBuilder) ForInit(operation string, handler interface{}) *actionformmodalBuilder {
	b.schema.XComponentProps["forInit"] = operation
	b.schema.HandlerNames = append(b.schema.HandlerNames, operation)
	controller.Bind(operation, handler)
	return b
}

func (b *actionformmodalBuilder) ForSubmit(operation string, handler interface{}) *actionformmodalBuilder {
	b.schema.XComponentProps["forSubmit"] = operation
	b.schema.HandlerNames = append(b.schema.HandlerNames, operation)
	controller.Bind(operation, handler)
	return b
}

func (b *actionformmodalBuilder) ButtonStyle(style map[string]interface{}) *actionformmodalBuilder {
	b.schema.XComponentProps["style"] = style
	return b
}

func (b *actionformmodalBuilder) ButtonType(typ string) *actionformmodalBuilder {
	b.schema.XComponentProps["type"] = typ
	return b
}

func (b *actionformmodalBuilder) ButtonStatus(status string) *actionformmodalBuilder {
	b.schema.XComponentProps["status"] = status
	return b
}

func (b *actionformmodalBuilder) ButtonSize(size string) *actionformmodalBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *actionformmodalBuilder) ButtonShape(shape string) *actionformmodalBuilder {
	b.schema.XComponentProps["shape"] = shape
	return b
}

func (b *actionformmodalBuilder) ButtonIcon(icon string) *actionformmodalBuilder {
	b.schema.XComponentProps["icon"] = icon
	return b
}

func (b *actionformmodalBuilder) ButtonIconOnly() *actionformmodalBuilder {
	b.schema.XComponentProps["iconOnly"] = true
	return b
}

func (b *actionformmodalBuilder) ButtonLong() *actionformmodalBuilder {
	b.schema.XComponentProps["long"] = true
	return b
}
