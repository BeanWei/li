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

func (b *timerangeBuilder) Required() *ui.Schema {
	b.schema.Required = true
	return b.schema
}

func (b *timerangeBuilder) Title(title string) *ui.Schema {
	b.schema.Title = title
	return b.schema
}

func (b *timerangeBuilder) Format(format string) *ui.Schema {
	b.schema.XComponentProps["format"] = format
	return b.schema
}
