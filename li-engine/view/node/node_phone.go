package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Phone(name string) *phoneBuilder {
	return &phoneBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeString,
			XComponent: ui.ComponentInput,
			XDecorator: ui.DecoratorFormItem,
			XValidator: "phone",
		},
	}}
}

type phoneBuilder struct {
	*NodeBuilder
}
