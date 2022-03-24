package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Email(name string) *emailBuilder {
	return &emailBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeString,
			XComponent: ui.ComponentInput,
			XDecorator: ui.DecoratorFormItem,
			XValidator: "email",
		},
	}}
}

type emailBuilder struct {
	*NodeBuilder
}
