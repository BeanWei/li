package node

import "github.com/BeanWei/li/li-engine/view/ui"

func RadioGroup(name string) *radiogroupBuilder {
	return &radiogroupBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			XPath:      name,
			Type:       ui.SchemaTypeAny,
			XComponent: ui.ComponentRadioGroup,
			XDecorator: ui.DecoratorFormItem,
			Enum:       make([]map[string]interface{}, 0),
		},
	}}
}

type radiogroupBuilder struct {
	*NodeBuilder
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

func (b *radiogroupBuilder) Option(label string, value ...interface{}) *radiogroupBuilder {
	var val interface{}
	if len(value) > 0 {
		val = value[0]
	} else {
		val = label
	}
	b.schema.Enum = append(b.schema.Enum, map[string]interface{}{
		"label": label,
		"value": val,
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
