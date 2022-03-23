package node

import (
	"github.com/BeanWei/li/li-engine/control"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func List(name string) *listBuilder {
	return &listBuilder{schema: &ui.Schema{
		Name:            name,
		Type:            ui.SchemaTypeVoid,
		XComponent:      ui.ComponentList,
		XComponentProps: make(map[string]interface{}),
		Properties:      make(map[string]*ui.Schema),
	}}
}

type listBuilder struct {
	schema *ui.Schema
}

func (b *listBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *listBuilder) DecoratorCard() *listBuilder {
	b.schema.XDecorator = ui.DecoratorCardItem
	return b
}

func (b *listBuilder) ForInit(operation string, controller interface{}) *listBuilder {
	b.schema.XComponentProps["forInit"] = operation
	control.RegisterController(operation, controller)
	return b
}

func (b *listBuilder) Children(elements ...view.Node) *listBuilder {
	for _, element := range elements {
		b.schema.Properties[element.Schema().Name] = element.Schema()
	}
	return b
}
