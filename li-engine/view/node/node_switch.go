package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Switch(name string) *switchBuilder {
	return &switchBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeBool,
			XComponent: ui.ComponentSwitch,
			XDecorator: ui.DecoratorFormItem,
		},
	}}
}

type switchBuilder struct {
	*NodeBuilder
}
