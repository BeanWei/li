package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func RadioGroup(name string) *radiogroupBuilder {
	return &radiogroupBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeAny,
			XComponent:      ui.ComponentRadioGroup,
			XComponentProps: make(map[string]interface{}),
			XDecorator:      ui.DecoratorFormItem,
			Enum:            make([]map[string]interface{}, 0),
		},
	}}
}

type radiogroupBuilder struct {
	*NodeBuilder
}

func (b *radiogroupBuilder) AC(f ac.AC) *radiogroupBuilder {
	b.schema.AC = f
	return b
}

func (b *radiogroupBuilder) Title(title string) *radiogroupBuilder {
	b.schema.Title = title
	return b
}

func (b *radiogroupBuilder) Description(description string) *radiogroupBuilder {
	b.schema.Description = description
	return b
}

func (b *radiogroupBuilder) Default(value interface{}) *radiogroupBuilder {
	b.schema.Default = value
	return b
}

func (b *radiogroupBuilder) Option(value interface{}, label ...string) *radiogroupBuilder {
	var label_ interface{}
	if len(label) > 0 {
		label_ = label[0]
	} else {
		label_ = value
	}
	b.schema.Enum = append(b.schema.Enum, map[string]interface{}{
		"label": label_,
		"value": value,
	})
	return b
}

func (b *radiogroupBuilder) Type(typ string) *radiogroupBuilder {
	b.schema.XComponentProps["type"] = typ
	return b
}

func (b *radiogroupBuilder) Direction(direction string) *radiogroupBuilder {
	b.schema.XComponentProps["direction"] = direction
	return b
}
