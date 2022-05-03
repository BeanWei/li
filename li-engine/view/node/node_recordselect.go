package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/controller"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func RecordSelect(name string) *recordselectBuilder {
	return &recordselectBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeAny,
			XComponent: ui.ComponentRecordSelect,
			XComponentProps: map[string]interface{}{
				"fieldNames": map[string]interface{}{
					"title": "id",
					"value": "id",
				},
				"searchConfig": map[string]interface{}{
					"operation": "",
				},
			},
			XDecorator: ui.DecoratorFormItem,
		},
	}}
}

type recordselectBuilder struct {
	*NodeBuilder
}

func (b *recordselectBuilder) AC(f ac.AC) *recordselectBuilder {
	b.schema.AC = f
	return b
}

func (b *recordselectBuilder) Title(title string) *recordselectBuilder {
	b.SetTitle(title)
	return b
}

func (b *recordselectBuilder) Description(description string) *recordselectBuilder {
	b.SetDescription(description)
	return b
}

func (b *recordselectBuilder) Default(value interface{}) *recordselectBuilder {
	b.schema.Default = value
	return b
}

func (b *recordselectBuilder) Multiple(multiple bool) *recordselectBuilder {
	b.schema.XComponentProps["multiple"] = multiple
	if multiple {
		b.schema.Type = ui.SchemaTypeArray
	}
	return b
}

func (b *recordselectBuilder) FieldNamesTitle(titleFieldName string) *recordselectBuilder {
	fn, ok := b.schema.XComponentProps["fieldNames"].(map[string]interface{})
	if ok {
		fn["title"] = titleFieldName
		b.schema.XComponentProps["fieldNames"] = fn
	}
	return b
}

func (b *recordselectBuilder) FieldNamesValue(valueFieldName string) *recordselectBuilder {
	fn, ok := b.schema.XComponentProps["fieldNames"].(map[string]interface{})
	if ok {
		fn["value"] = valueFieldName
		b.schema.XComponentProps["fieldNames"] = fn
	}
	return b
}

func (b *recordselectBuilder) FieldNamesAvatar(avatarFieldName string) *recordselectBuilder {
	fn, ok := b.schema.XComponentProps["fieldNames"].(map[string]interface{})
	if ok {
		fn["avatar"] = avatarFieldName
		b.schema.XComponentProps["fieldNames"] = fn
	}
	return b
}

func (b *recordselectBuilder) FieldNamesDescription(descriptionFieldName ...string) *recordselectBuilder {
	fn, ok := b.schema.XComponentProps["fieldNames"].(map[string]interface{})
	if ok {
		fn["description"] = descriptionFieldName
		b.schema.XComponentProps["fieldNames"] = fn
	}
	return b
}

func (b *recordselectBuilder) Size(size string) *recordselectBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *recordselectBuilder) MaxTagCount(max int) *recordselectBuilder {
	b.schema.XComponentProps["maxTagCount"] = max
	return b
}

func (b *recordselectBuilder) Prefix(prefix string) *recordselectBuilder {
	b.schema.XComponentProps["prefix"] = prefix
	return b
}

func (b *recordselectBuilder) SearchOperation(operation string, handler interface{}) *recordselectBuilder {
	fn, ok := b.schema.XComponentProps["searchConfig"].(map[string]interface{})
	if ok {
		fn["operation"] = operation
		b.schema.XComponentProps["searchConfig"] = fn
		b.schema.HandlerNames = append(b.schema.HandlerNames, operation)
		controller.Bind(operation, handler)
	}
	return b
}

func (b *recordselectBuilder) SearchVariables(variables string) *recordselectBuilder {
	sc, ok := b.schema.XComponentProps["searchConfig"].(map[string]interface{})
	if ok {
		sc["variables"] = variables
		b.schema.XComponentProps["searchConfig"] = sc
	}
	return b
}
