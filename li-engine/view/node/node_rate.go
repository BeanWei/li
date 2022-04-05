package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func Rate(name string) *rateBuilder {
	return &rateBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeNumber,
			XComponent:      ui.ComponentRate,
			XComponentProps: make(map[string]interface{}),
			XDecorator:      ui.DecoratorFormItem,
		},
	}}
}

type rateBuilder struct {
	*NodeBuilder
}

func (b *rateBuilder) AC(f ac.AC) *rateBuilder {
	b.schema.AC = f
	return b
}

func (b *rateBuilder) Title(title string) *rateBuilder {
	b.schema.Title = title
	return b
}

func (b *rateBuilder) Description(description string) *rateBuilder {
	b.schema.Description = description
	return b
}

func (b *rateBuilder) Default(value interface{}) *rateBuilder {
	b.schema.Default = value
	return b
}

func (b *rateBuilder) Count(count int) *rateBuilder {
	b.schema.XComponentProps["count"] = count
	return b
}

func (b *rateBuilder) AllowHalf(allowHalf bool) *rateBuilder {
	b.schema.XComponentProps["allowHalf"] = allowHalf
	return b
}

func (b *rateBuilder) Grading(grading bool) *rateBuilder {
	b.schema.XComponentProps["grading"] = grading
	return b
}
