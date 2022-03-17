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

func (b *timeBuilder) Required() *timeBuilder {
	b.schema.Required = true
	return b
}

func (b *timeBuilder) Title(title string) *timeBuilder {
	b.schema.Title = title
	return b
}

func (b *timeBuilder) Format(format string) *timeBuilder {
	b.schema.XComponentProps["format"] = format
	return b
}
