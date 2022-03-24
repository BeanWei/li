package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Text(name string) *textBuilder {
	return &textBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeString,
			XComponent: ui.ComponentInput,
			XDecorator: ui.DecoratorFormItem,
		},
	}}
}

type textBuilder struct {
	*NodeBuilder
}
