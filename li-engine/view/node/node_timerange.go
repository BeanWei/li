package node

import "github.com/BeanWei/li/li-engine/view/ui"

func TimeRange(name string) *timerangeBuilder {
	return &timerangeBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeArray,
			XComponent:      ui.ComponentDatePickerRangePicker,
			XComponentProps: make(map[string]interface{}),
			XDecorator:      ui.DecoratorFormItem,
		},
	}}
}

type timerangeBuilder struct {
	*NodeBuilder
}

func (b *timerangeBuilder) Format(format string) *timerangeBuilder {
	b.schema.XComponentProps["format"] = format
	return b
}
