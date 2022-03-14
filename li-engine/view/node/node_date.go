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

func (b *dateBuilder) Required() *ui.Schema {
	b.schema.Required = true
	return b.schema
}

func (b *dateBuilder) Title(title string) *ui.Schema {
	b.schema.Title = title
	return b.schema
}

// Mode Time, Week, Month, Quarter, Year
func (b *dateBuilder) Mode(mode string) *ui.Schema {
	b.schema.XComponentProps["mode"] = mode
	return b.schema
}

func (b *dateBuilder) Format(format string) *ui.Schema {
	b.schema.XComponentProps["format"] = format
	return b.schema
}
