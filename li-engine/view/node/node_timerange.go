package node

import "github.com/BeanWei/li/li-engine/view/ui"

func TimeRange(name string) *timerangeBuilder {
	return &timerangeBuilder{schema: &ui.Schema{
		Name:            name,
		Type:            ui.SchemaTypeArray,
		XComponent:      ui.ComponentDatePickerRangePicker,
		XComponentProps: make(map[string]interface{}),
		XDecorator:      ui.DecoratorFormItem,
	}}
}

type timerangeBuilder struct {
	schema *ui.Schema
}

func (b *timerangeBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *timerangeBuilder) Required() *timerangeBuilder {
	b.schema.Required = true
	return b
}

func (b *timerangeBuilder) Title(title string) *timerangeBuilder {
	b.schema.Title = title
	return b
}

func (b *timerangeBuilder) Format(format string) *timerangeBuilder {
	b.schema.XComponentProps["format"] = format
	return b
}
