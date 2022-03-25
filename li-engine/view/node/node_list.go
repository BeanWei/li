package node

import (
	"github.com/BeanWei/li/li-engine/controller"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func List(name string) *listBuilder {
	return &listBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentList,
			XComponentProps: make(map[string]interface{}),
			Properties:      make(map[string]*ui.Schema),
		},
	}}
}

type listBuilder struct {
	*NodeBuilder
}

func (b *listBuilder) DecoratorCard() *listBuilder {
	b.schema.XDecorator = ui.DecoratorCardItem
	return b
}

func (b *listBuilder) ForInit(operation string, handler interface{}) *listBuilder {
	b.schema.XComponentProps["forInit"] = operation
	controller.Bind(operation, handler)
	return b
}
