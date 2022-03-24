package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Time(name string) *timeBuilder {
	return &timeBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeString,
			XComponent:      ui.ComponentTimePicker,
			XComponentProps: make(map[string]interface{}),
			XDecorator:      ui.DecoratorFormItem,
		},
	}}
}

type timeBuilder struct {
	*NodeBuilder
}

func (b *timeBuilder) Format(format string) *timeBuilder {
	b.schema.XComponentProps["format"] = format
	return b
}
