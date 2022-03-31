package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func CheckboxGroup(name string) *checkboxgroupBuilder {
	return &checkboxgroupBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeArray,
			XComponent:      ui.ComponentCheckboxGroup,
			XComponentProps: make(map[string]interface{}),
			XDecorator:      ui.DecoratorFormItem,
			Enum:            make([]map[string]interface{}, 0),
		},
	}}
}

type checkboxgroupBuilder struct {
	*NodeBuilder
}

func (b *checkboxgroupBuilder) AC(f ac.AC) *checkboxgroupBuilder {
	b.schema.AC = f
	return b
}

func (b *checkboxgroupBuilder) Title(title string) *checkboxgroupBuilder {
	b.schema.Title = title
	return b
}

func (b *checkboxgroupBuilder) Description(description string) *checkboxgroupBuilder {
	b.schema.Description = description
	return b
}

func (b *checkboxgroupBuilder) Default(value interface{}) *checkboxgroupBuilder {
	b.schema.Default = value
	return b
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

func (b *checkboxgroupBuilder) Direction(direction string) *checkboxgroupBuilder {
	b.schema.XComponentProps["direction"] = direction
	return b
}
