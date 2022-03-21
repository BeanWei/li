package node

import (
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func Form(name string) *formBuilder {
	return &formBuilder{schema: &ui.Schema{
		Name:       name,
		Type:       ui.SchemaTypeObject,
		XComponent: ui.ComponentForm,
	}}
}

type formBuilder struct {
	schema *ui.Schema
}

func (b *formBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *formBuilder) DecoratorCard() *formBuilder {
	b.schema.XDecorator = ui.DecoratorCardItem
	return b
}

func (b *formBuilder) Children(elements ...view.Node) *formBuilder {
	for _, element := range elements {
		b.schema.Properties[element.Schema().Name] = element.Schema()
	}
	return b
}
