package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Password(name string) *passwordBuilder {
	return &passwordBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeString,
			XComponent: ui.ComponentPassword,
			XDecorator: ui.DecoratorFormItem,
		},
	}}
}

type passwordBuilder struct {
	*NodeBuilder
}
