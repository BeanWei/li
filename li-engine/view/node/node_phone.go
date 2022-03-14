package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Phone(name string) *phoneBuilder {
	return &phoneBuilder{schema: &ui.Schema{
		Name:       name,
		Type:       ui.SchemaTypeString,
		XComponent: ui.ComponentInput,
		XDecorator: ui.DecoratorFormItem,
		XValidator: "phone",
	}}
}

type phoneBuilder struct {
	schema *ui.Schema
}

func (b *phoneBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *phoneBuilder) Required() *ui.Schema {
	b.schema.Required = true
	return b.schema
}

func (b *phoneBuilder) Title(title string) *ui.Schema {
	b.schema.Title = title
	return b.schema
}
