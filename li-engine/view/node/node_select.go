package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func Select(name string) *selectBuilder {
	return &selectBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeAny,
			XComponent:      ui.ComponentSelect,
			XDecorator:      ui.DecoratorFormItem,
			XComponentProps: make(map[string]interface{}),
			Enum:            make([]map[string]interface{}, 0),
		},
	}}
}

type selectBuilder struct {
	*NodeBuilder
}

func (b *selectBuilder) AC(f ac.AC) *selectBuilder {
	b.schema.AC = f
	return b
}

func (b *selectBuilder) Title(title string) *selectBuilder {
	b.schema.Title = title
	return b
}

func (b *selectBuilder) Description(description string) *selectBuilder {
	b.schema.Description = description
	return b
}

func (b *selectBuilder) Default(value interface{}) *selectBuilder {
	b.schema.Default = value
	return b
}

func (b *selectBuilder) Multiple() *selectBuilder {
	b.schema.XComponentProps["mode"] = "multiple"
	return b
}

func (b *selectBuilder) Size(size string) *selectBuilder {
	b.schema.XComponentProps["size"] = size
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

func (b *selectBuilder) Prefix(prefix string) *selectBuilder {
	b.schema.XComponentProps["prefix"] = prefix
	return b
}

func (b *selectBuilder) DragToSort() *selectBuilder {
	b.schema.XComponentProps["dragToSort"] = true
	return b
}

func (b *selectBuilder) Option(value interface{}, label ...string) *selectBuilder {
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
