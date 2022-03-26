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
		},
	}}
}

type submitBuilder struct {
	*NodeBuilder
}

func (b *submitBuilder) ForSubmit(operation string, handler interface{}) *submitBuilder {
	b.schema.XComponentProps["forSubmit"] = operation
	controller.Bind(operation, handler)
	return b
}

func (b *submitBuilder) ForSubmitSuccessTo(to string) *submitBuilder {
	b.schema.XComponentProps["forSubmitSuccessTo"] = to
	return b
}

func (b *submitBuilder) ButtonLong() *submitBuilder {
	b.schema.XComponentProps["long"] = true
	return b
}

func (b *submitBuilder) ButtonType(typ string) *submitBuilder {
	b.schema.XComponentProps["type"] = typ
	return b
}

func (b *submitBuilder) ButtonStyle(style map[string]interface{}) *submitBuilder {
	b.schema.XComponentProps["style"] = style
	return b
}
