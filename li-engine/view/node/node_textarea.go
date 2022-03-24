package node

import "github.com/BeanWei/li/li-engine/view/ui"

func TextArea(name string) *textareaBuilder {
	return &textareaBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeString,
			XComponent: ui.ComponentInputTextArea,
			XDecorator: ui.DecoratorFormItem,
		},
	}}
}

type textareaBuilder struct {
	*NodeBuilder
}
