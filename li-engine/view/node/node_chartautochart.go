package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/controller"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ChartAutoChart(name string) *chartautochartBuilder {
	return &chartautochartBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentChartAutoChart,
			XComponentProps: make(map[string]interface{}),
			XDecorator:      ui.DecoratorChartItem,
			XDecoratorProps: make(map[string]interface{}),
		},
	}}
}

type chartautochartBuilder struct {
	*NodeBuilder
}

func (b *chartautochartBuilder) AC(f ac.AC) *chartautochartBuilder {
	b.schema.AC = f
	return b
}

func (b *chartautochartBuilder) Title(title string) *chartautochartBuilder {
	b.schema.XDecoratorProps["title"] = title
	return b
}

func (b *chartautochartBuilder) SubTitle(subTitle string) *chartautochartBuilder {
	b.schema.XDecoratorProps["subTitle"] = subTitle
	return b
}

func (b *chartautochartBuilder) ForInitVariables(variables map[string]interface{}) *chartautochartBuilder {
	b.schema.XDecoratorProps["forInitVariables"] = variables
	return b
}

func (b *chartautochartBuilder) Description(description string) *chartautochartBuilder {
	b.SetDescription(description)
	return b
}

func (b *chartautochartBuilder) ForInit(operation string, handler interface{}) *chartautochartBuilder {
	b.schema.XDecoratorProps["forInit"] = operation
	b.schema.HandlerNames = append(b.schema.HandlerNames, operation)
	controller.Bind(operation, handler)
	return b
}

func (b *chartautochartBuilder) GridSpan(gridSpan int) *chartautochartBuilder {
	b.schema.XDecoratorProps["gridSpan"] = gridSpan
	return b
}

func (b *chartautochartBuilder) Style(style map[string]interface{}) *chartautochartBuilder {
	b.schema.XComponentProps["style"] = style
	return b
}

func (b *chartautochartBuilder) Fields(fields ...string) *chartautochartBuilder {
	b.schema.XComponentProps["fields"] = fields
	return b
}

func (b *chartautochartBuilder) SmartColor(smartColor bool) *chartautochartBuilder {
	b.schema.XComponentProps["smartColor"] = smartColor
	return b
}

func (b *chartautochartBuilder) OptionPurpose(purpose string) *chartautochartBuilder {
	options, ok := b.schema.XComponentProps["options"].(map[string]interface{})
	if ok {
		options["purpose"] = purpose
		b.schema.XComponentProps["options"] = options
	}
	return b
}

func (b *chartautochartBuilder) OptionRefine(refine bool) *chartautochartBuilder {
	options, ok := b.schema.XComponentProps["options"].(map[string]interface{})
	if ok {
		options["refine"] = refine
		b.schema.XComponentProps["options"] = options
	}
	return b
}

func (b *chartautochartBuilder) OptionFields(fields ...string) *chartautochartBuilder {
	options, ok := b.schema.XComponentProps["options"].(map[string]interface{})
	if ok {
		options["fields"] = fields
		b.schema.XComponentProps["options"] = options
	}
	return b
}
