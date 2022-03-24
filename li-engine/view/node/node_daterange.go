package node

import "github.com/BeanWei/li/li-engine/view/ui"

func DateRange(name string) *daterangeBuilder {
	return &daterangeBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeArray,
			XComponent:      ui.ComponentDatePickerRangePicker,
			XComponentProps: make(map[string]interface{}),
			XDecorator:      ui.DecoratorFormItem,
		},
	}}
}

type daterangeBuilder struct {
	*NodeBuilder
}

// Mode Time, Week, Month, Quarter, Year
func (b *daterangeBuilder) Mode(mode string) *daterangeBuilder {
	b.schema.XComponentProps["mode"] = mode
	return b
}

func (b *daterangeBuilder) Format(format string) *daterangeBuilder {
	b.schema.XComponentProps["format"] = format
	return b
}
