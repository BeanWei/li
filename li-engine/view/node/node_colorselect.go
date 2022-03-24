package node

import "github.com/BeanWei/li/li-engine/view/ui"

func ColorSelect(name string) *colorselectBuilder {
	return &colorselectBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeString,
			XComponent: ui.ComponentColorSelect,
			XDecorator: ui.DecoratorFormItem,
		},
	}}
}

type colorselectBuilder struct {
	*NodeBuilder
}
