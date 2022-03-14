package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Rate(name string) *rateBuilder {
	return &rateBuilder{schema: &ui.Schema{
		Name:            name,
		Type:            ui.SchemaTypeNumber,
		XComponent:      ui.ComponentRate,
		XComponentProps: make(map[string]interface{}),
		XDecorator:      ui.DecoratorFormItem,
	}}
}

type rateBuilder struct {
	schema *ui.Schema
}

func (b *rateBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *rateBuilder) Required() *ui.Schema {
	b.schema.Required = true
	return b.schema
}

func (b *rateBuilder) Title(title string) *ui.Schema {
	b.schema.Title = title
	return b.schema
}

func (b *rateBuilder) Count(count int) *ui.Schema {
	b.schema.XComponentProps["count"] = count
	return b.schema
}

func (b *rateBuilder) AllowHalf() *ui.Schema {
	b.schema.XComponentProps["allowHalf"] = true
	return b.schema
}

func (b *rateBuilder) Grading() *ui.Schema {
	b.schema.XComponentProps["grading"] = true
	return b.schema
}
