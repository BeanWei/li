package node

import "github.com/BeanWei/li/li-engine/view/ui"

func CheckboxGroup(name string) *checkboxgroupBuilder {
	return &checkboxgroupBuilder{schema: &ui.Schema{
		Name:       name,
		Type:       ui.SchemaTypeArray,
		XComponent: ui.ComponentCheckboxGroup,
		XDecorator: ui.DecoratorFormItem,
		Enum:       make([]map[string]interface{}, 0),
	}}
}

type checkboxgroupBuilder struct {
	schema *ui.Schema
}

func (b *checkboxgroupBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *checkboxgroupBuilder) Required() *ui.Schema {
	b.schema.Required = true
	return b.schema
}

func (b *checkboxgroupBuilder) Title(title string) *ui.Schema {
	b.schema.Title = title
	return b.schema
}

func (b *checkboxgroupBuilder) Option(label string, value ...interface{}) *ui.Schema {
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
	return b.schema
}
