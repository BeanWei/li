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

func (b *rateBuilder) Required() *rateBuilder {
	b.schema.Required = true
	return b
}

func (b *rateBuilder) Title(title string) *rateBuilder {
	b.schema.Title = title
	return b
}

func (b *rateBuilder) Count(count int) *rateBuilder {
	b.schema.XComponentProps["count"] = count
	return b
}

func (b *rateBuilder) AllowHalf() *rateBuilder {
	b.schema.XComponentProps["allowHalf"] = true
	return b
}

func (b *rateBuilder) Grading() *rateBuilder {
	b.schema.XComponentProps["grading"] = true
	return b
}
