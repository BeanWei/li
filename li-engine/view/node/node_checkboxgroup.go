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

func (b *checkboxgroupBuilder) Option(value interface{}, label ...string) *checkboxgroupBuilder {
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

func (b *checkboxgroupBuilder) Direction(direction string) *checkboxgroupBuilder {
	b.schema.XComponentProps["direction"] = direction
	return b
}
