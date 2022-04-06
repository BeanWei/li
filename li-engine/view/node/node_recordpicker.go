package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func RecordPicker(name string) *recordpickerBuilder {
	return &recordpickerBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeAny,
			XComponent: ui.ComponentRecordPicker,
			XComponentProps: map[string]interface{}{
				"fieldNames": map[string]string{
					"value": "id",
					"label": "id",
				},
			},
			XDecorator: ui.DecoratorFormItem,
		},
	}}
}

type recordpickerBuilder struct {
	*NodeBuilder
}

func (b *recordpickerBuilder) AC(f ac.AC) *recordpickerBuilder {
	b.schema.AC = f
	return b
}

func (b *recordpickerBuilder) Title(title string) *recordpickerBuilder {
	b.schema.Title = title
	return b
}

func (b *recordpickerBuilder) Description(description string) *recordpickerBuilder {
	b.schema.Description = description
	return b
}

func (b *recordpickerBuilder) Default(value interface{}) *recordpickerBuilder {
	b.schema.Default = value
	return b
}

func (b *recordpickerBuilder) SelectionMultiple(multiple bool) *recordpickerBuilder {
	b.schema.XComponentProps["multiple"] = multiple
	return b
}

func (b *recordpickerBuilder) FieldNamesLabel(labelFieldName string) *recordpickerBuilder {
	fn, ok := b.schema.XComponentProps["fieldNames"].(map[string]string)
	if ok {
		fn["label"] = labelFieldName
		b.schema.XComponentProps["fieldNames"] = fn
	}
	return b
}

func (b *recordpickerBuilder) FieldNamesValue(valueFieldName string) *recordpickerBuilder {
	fn, ok := b.schema.XComponentProps["fieldNames"].(map[string]string)
	if ok {
		fn["value"] = valueFieldName
		b.schema.XComponentProps["fieldNames"] = fn
	}
	return b
}

func (b *recordpickerBuilder) Size(size string) *recordpickerBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *recordpickerBuilder) MaxTagCount(max int) *recordpickerBuilder {
	b.schema.XComponentProps["maxTagCount"] = max
	return b
}

func (b *recordpickerBuilder) Prefix(prefix string) *recordpickerBuilder {
	b.schema.XComponentProps["prefix"] = prefix
	return b
}
