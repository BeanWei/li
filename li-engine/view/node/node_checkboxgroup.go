package node

import "github.com/BeanWei/li/li-engine/view/ui"

func CheckboxGroup(name string) *checkboxgroupBuilder {
	return &checkboxgroupBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeArray,
			XComponent: ui.ComponentCheckboxGroup,
			XDecorator: ui.DecoratorFormItem,
			Enum:       make([]map[string]interface{}, 0),
		},
	}}
}

type checkboxgroupBuilder struct {
	*NodeBuilder
}

func (b *checkboxgroupBuilder) Option(label string, value ...interface{}) *checkboxgroupBuilder {
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
