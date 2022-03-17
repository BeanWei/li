package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Select(name string) *selectBuilder {
	return &selectBuilder{schema: &ui.Schema{
		Name:            name,
		Type:            ui.SchemaTypeAny,
		XComponent:      ui.ComponentSelect,
		XDecorator:      ui.DecoratorFormItem,
		XComponentProps: make(map[string]interface{}),
		Enum:            make([]map[string]interface{}, 0),
	}}
}

type selectBuilder struct {
	schema *ui.Schema
}

func (b *selectBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *selectBuilder) Required() *selectBuilder {
	b.schema.Required = true
	return b
}

func (b *selectBuilder) Title(title string) *selectBuilder {
	b.schema.Title = title
	return b
}

func (b *selectBuilder) Multiple() *selectBuilder {
	b.schema.XComponentProps["mode"] = "multiple"
	return b
}

func (b *selectBuilder) AllowCreate() *selectBuilder {
	b.schema.XComponentProps["allowCreate"] = true
	return b
}

func (b *selectBuilder) MaxTagCount(max int) *selectBuilder {
	b.schema.XComponentProps["maxTagCount"] = max
	return b
}

func (b *selectBuilder) Option(label string, value ...interface{}) *selectBuilder {
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
