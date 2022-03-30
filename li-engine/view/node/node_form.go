package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Form(name string) *formBuilder {
	return &formBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			XPath:      name,
			Type:       ui.SchemaTypeObject,
			XComponent: ui.ComponentForm,
		},
	}}
}

type formBuilder struct {
	*NodeBuilder
}

func (b *formBuilder) DecoratorCard() *formBuilder {
	b.schema.XDecorator = ui.DecoratorCardItem
	return b
}
