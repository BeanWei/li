package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Checkbox(name string) *checkboxBuilder {
	return &checkboxBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeBool,
			XComponent: ui.ComponentCheckbox,
			XDecorator: ui.DecoratorFormItem,
		},
	}}
}

type checkboxBuilder struct {
	*NodeBuilder
}
