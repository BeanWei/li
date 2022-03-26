package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Password(name string) *passwordBuilder {
	return &passwordBuilder{&textBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeString,
			XComponent:      ui.ComponentPassword,
			XComponentProps: make(map[string]interface{}),
			XDecorator:      ui.DecoratorFormItem,
		},
	}}}
}

type passwordBuilder struct {
	*textBuilder
}
