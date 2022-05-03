package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/controller"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func FlowEditor(name string) *floweditorBuilder {
	return &floweditorBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeString,
			XComponent: ui.ComponentFlowEditor,
			XComponentProps: map[string]interface{}{
				"fetchUserConfig": map[string]interface{}{
					"operation": "",
					"fieldNames": map[string]string{
						"label": "id",
						"value": "id",
					},
				},
			},
			XDecorator: ui.DecoratorFormItem,
		},
	}}
}

type floweditorBuilder struct {
	*NodeBuilder
}

func (b *floweditorBuilder) AC(f ac.AC) *floweditorBuilder {
	b.schema.AC = f
	return b
}

func (b *floweditorBuilder) Title(title string) *floweditorBuilder {
	b.SetTitle(title)
	return b
}

func (b *floweditorBuilder) Description(description string) *floweditorBuilder {
	b.SetDescription(description)
	return b
}

func (b *floweditorBuilder) Default(value interface{}) *floweditorBuilder {
	b.schema.Default = value
	return b
}

func (b *recordselectBuilder) Height(height string) *recordselectBuilder {
	b.schema.XComponentProps["height"] = height
	return b
}

func (b *floweditorBuilder) FetchUserConfigOperation(operation string, handler interface{}) *floweditorBuilder {
	fn, ok := b.schema.XComponentProps["fetchUserConfig"].(map[string]interface{})
	if ok {
		fn["operation"] = operation
		b.schema.XComponentProps["fetchUserConfig"] = fn
		b.schema.HandlerNames = append(b.schema.HandlerNames, operation)
		controller.Bind(operation, handler)
	}
	return b
}

func (b *recordselectBuilder) FetchUserConfigVariables(variables string) *recordselectBuilder {
	sc, ok := b.schema.XComponentProps["fetchUserConfig"].(map[string]interface{})
	if ok {
		sc["variables"] = variables
		b.schema.XComponentProps["fetchUserConfig"] = sc
	}
	return b
}

func (b *floweditorBuilder) FetchUserConfigFieldNamesLabel(labelFieldName string) *floweditorBuilder {
	fn, ok := b.schema.XComponentProps["fetchUserConfig"].(map[string]interface{})
	if ok {
		fn2, ok2 := fn["fieldNames"].(map[string]string)
		if ok2 {
			fn2["label"] = labelFieldName
		}
		fn["fieldNames"] = fn2
		b.schema.XComponentProps["fetchUserConfig"] = fn
	}
	return b
}

func (b *floweditorBuilder) FetchUserConfigFieldNamesValue(valueFieldName string) *floweditorBuilder {
	fn, ok := b.schema.XComponentProps["fetchUserConfig"].(map[string]interface{})
	if ok {
		fn2, ok2 := fn["fieldNames"].(map[string]string)
		if ok2 {
			fn2["value"] = valueFieldName
		}
		fn["fieldNames"] = fn2
		b.schema.XComponentProps["fetchUserConfig"] = fn
	}
	return b
}
