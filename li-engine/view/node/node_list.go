package node

import (
	"github.com/BeanWei/li/li-engine/controller"
	"github.com/BeanWei/li/li-engine/view/ui"
	"github.com/gogf/gf/v2/container/gmap"
)

func List(name string) *listBuilder {
	return &listBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentList,
			XComponentProps: make(map[string]interface{}),
			Properties:      gmap.NewListMap(),
		},
	}}
}

type listBuilder struct {
	*NodeBuilder
}

func (b *listBuilder) Title(title string) *listBuilder {
	b.schema.Title = title
	return b
}

func (b *listBuilder) Description(description string) *listBuilder {
	b.schema.Description = description
	return b
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
