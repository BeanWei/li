package node

import (
	"github.com/BeanWei/li/li-engine/controller"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func Submit(name string) *submitBuilder {
	return &submitBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentSubmit,
			XDecorator:      ui.DecoratorFormItem,
			XComponentProps: make(map[string]interface{}),
			HandlerNames:    make([]string, 0),
		},
	}}
}

type submitBuilder struct {
	*NodeBuilder
}

func (b *submitBuilder) Title(title string) *submitBuilder {
	b.schema.Title = title
	return b
}

func (b *submitBuilder) Description(description string) *submitBuilder {
	b.schema.Description = description
	return b
}

func (b *submitBuilder) ForSubmit(operation string, handler interface{}) *submitBuilder {
	b.schema.XComponentProps["forSubmit"] = operation
	b.schema.HandlerNames = append(b.schema.HandlerNames, operation)
	controller.Bind(operation, handler)
	return b
}

func (b *submitBuilder) ForSubmitSuccessTo(to string) *submitBuilder {
	b.schema.XComponentProps["forSubmitSuccessTo"] = to
	return b
}

func (b *submitBuilder) ButtonStyle(style map[string]interface{}) *submitBuilder {
	b.schema.XComponentProps["style"] = style
	return b
}

func (b *submitBuilder) ButtonType(typ string) *submitBuilder {
	b.schema.XComponentProps["type"] = typ
	return b
}

func (b *submitBuilder) ButtonStatus(status string) *submitBuilder {
	b.schema.XComponentProps["status"] = status
	return b
}

func (b *submitBuilder) ButtonSize(size string) *submitBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *submitBuilder) ButtonShape(shape string) *submitBuilder {
	b.schema.XComponentProps["shape"] = shape
	return b
}

func (b *submitBuilder) ButtonIcon(icon string) *submitBuilder {
	b.schema.XComponentProps["icon"] = icon
	return b
}

func (b *submitBuilder) ButtonIconOnly() *submitBuilder {
	b.schema.XComponentProps["iconOnly"] = true
	return b
}

func (b *submitBuilder) ButtonLong() *submitBuilder {
	b.schema.XComponentProps["long"] = true
	return b
}
