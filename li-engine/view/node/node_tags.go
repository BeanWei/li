package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Tags(name string) *tagsBuilder {
	return &tagsBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeArray,
			XComponent: ui.ComponentInputTag,
			XDecorator: ui.DecoratorFormItem,
		},
	}}
}

type tagsBuilder struct {
	*NodeBuilder
}
