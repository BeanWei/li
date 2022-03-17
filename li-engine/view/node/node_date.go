package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Date(name string) *dateBuilder {
	return &dateBuilder{schema: &ui.Schema{
		Name:            name,
		Type:            ui.SchemaTypeString,
		XComponent:      ui.ComponentDatePicker,
		XComponentProps: make(map[string]interface{}),
		XDecorator:      ui.DecoratorFormItem,
	}}
}

type dateBuilder struct {
	schema *ui.Schema
}

func (b *dateBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *dateBuilder) Required() *dateBuilder {
	b.schema.Required = true
	return b
}

func (b *dateBuilder) Title(title string) *dateBuilder {
	b.schema.Title = title
	return b
}

// Mode Time, Week, Month, Quarter, Year
func (b *dateBuilder) Mode(mode string) *dateBuilder {
	b.schema.XComponentProps["mode"] = mode
	return b
}

func (b *dateBuilder) Format(format string) *dateBuilder {
	b.schema.XComponentProps["format"] = format
	return b
}
