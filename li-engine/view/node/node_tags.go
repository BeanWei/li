package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Tags(name string) *tagsBuilder {
	return &tagsBuilder{schema: &ui.Schema{
		Name:       name,
		Type:       ui.SchemaTypeArray,
		XComponent: ui.ComponentInputTag,
		XDecorator: ui.DecoratorFormItem,
	}}
}

type tagsBuilder struct {
	schema *ui.Schema
}

func (b *tagsBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *tagsBuilder) Required() *tagsBuilder {
	b.schema.Required = true
	return b
}

func (b *tagsBuilder) Title(title string) *tagsBuilder {
	b.schema.Title = title
	return b
}
