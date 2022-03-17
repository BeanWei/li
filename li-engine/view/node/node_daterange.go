package node

import "github.com/BeanWei/li/li-engine/view/ui"

func DateRange(name string) *daterangeBuilder {
	return &daterangeBuilder{schema: &ui.Schema{
		Name:            name,
		Type:            ui.SchemaTypeArray,
		XComponent:      ui.ComponentDatePickerRangePicker,
		XComponentProps: make(map[string]interface{}),
		XDecorator:      ui.DecoratorFormItem,
	}}
}

type daterangeBuilder struct {
	schema *ui.Schema
}

func (b *daterangeBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *daterangeBuilder) Required() *daterangeBuilder {
	b.schema.Required = true
	return b
}

func (b *daterangeBuilder) Title(title string) *daterangeBuilder {
	b.schema.Title = title
	return b
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
