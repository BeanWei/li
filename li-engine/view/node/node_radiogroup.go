package node

import "github.com/BeanWei/li/li-engine/view/ui"

func RadioGroup(name string) *radiogroupBuilder {
	return &radiogroupBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
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
