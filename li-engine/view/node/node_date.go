package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Date(name string) *dateBuilder {
	return &dateBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeString,
			XComponent:      ui.ComponentDatePicker,
			XComponentProps: make(map[string]interface{}),
			XDecorator:      ui.DecoratorFormItem,
		},
	}}
}

type dateBuilder struct {
	*NodeBuilder
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
