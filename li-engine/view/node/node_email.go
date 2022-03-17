package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Email(name string) *emailBuilder {
	return &emailBuilder{schema: &ui.Schema{
		Name:       name,
		Type:       ui.SchemaTypeString,
		XComponent: ui.ComponentInput,
		XDecorator: ui.DecoratorFormItem,
		XValidator: "email",
	}}
}

type emailBuilder struct {
	schema *ui.Schema
}

func (b *emailBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *emailBuilder) Required() *emailBuilder {
	b.schema.Required = true
	return b
}

func (b *emailBuilder) Title(title string) *emailBuilder {
	b.schema.Title = title
	return b
}
