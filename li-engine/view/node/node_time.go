package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Time(name string) *timeBuilder {
	return &timeBuilder{schema: &ui.Schema{
		Name:            name,
		Type:            ui.SchemaTypeString,
		XComponent:      ui.ComponentTimePicker,
		XComponentProps: make(map[string]interface{}),
		XDecorator:      ui.DecoratorFormItem,
	}}
}

type timeBuilder struct {
	schema *ui.Schema
}

func (b *timeBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *timeBuilder) Required() *ui.Schema {
	b.schema.Required = true
	return b.schema
}

func (b *timeBuilder) Title(title string) *ui.Schema {
	b.schema.Title = title
	return b.schema
}

func (b *timeBuilder) Format(format string) *ui.Schema {
	b.schema.XComponentProps["format"] = format
	return b.schema
}
