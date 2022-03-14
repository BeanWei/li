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

func (b *selectBuilder) Required() *ui.Schema {
	b.schema.Required = true
	return b.schema
}

func (b *selectBuilder) Title(title string) *ui.Schema {
	b.schema.Title = title
	return b.schema
}

func (b *selectBuilder) Multiple() *ui.Schema {
	b.schema.XComponentProps["mode"] = "multiple"
	return b.schema
}

func (b *selectBuilder) AllowCreate() *ui.Schema {
	b.schema.XComponentProps["allowCreate"] = true
	return b.schema
}

func (b *selectBuilder) MaxTagCount(max int) *ui.Schema {
	b.schema.XComponentProps["maxTagCount"] = max
	return b.schema
}

func (b *selectBuilder) Option(label string, value ...interface{}) *ui.Schema {
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
