package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/controller"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ChartPie(name string) *chartpieBuilder {
	return &chartpieBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeVoid,
			XComponent: ui.ComponentChartPie,
			XComponentProps: map[string]interface{}{
				"meta": make([]map[string]interface{}, 0),
			},
			XDecorator:      ui.DecoratorChartItem,
			XDecoratorProps: make(map[string]interface{}),
		},
	}}
}

type chartpieBuilder struct {
	*NodeBuilder
}

func (b *chartpieBuilder) AC(f ac.AC) *chartpieBuilder {
	b.schema.AC = f
	return b
}

func (b *chartpieBuilder) Title(title string) *chartpieBuilder {
	b.schema.XDecoratorProps["title"] = title
	return b
}

func (b *chartpieBuilder) SubTitle(subTitle string) *chartpieBuilder {
	b.schema.XDecoratorProps["subTitle"] = subTitle
	return b
}

func (b *chartpieBuilder) ForInitVariables(variables map[string]interface{}) *chartpieBuilder {
	b.schema.XDecoratorProps["forInitVariables"] = variables
	return b
}

func (b *chartpieBuilder) Description(description string) *chartpieBuilder {
	b.schema.Description = description
	return b
}

func (b *chartpieBuilder) ForInit(operation string, handler interface{}) *chartpieBuilder {
	b.schema.XDecoratorProps["forInit"] = operation
	b.schema.HandlerNames = append(b.schema.HandlerNames, operation)
	controller.Bind(operation, handler)
	return b
}

func (b *chartpieBuilder) GridSpan(gridSpan int) *chartpieBuilder {
	b.schema.XDecoratorProps["gridSpan"] = gridSpan
	return b
}

func (b *chartpieBuilder) Style(style map[string]interface{}) *chartpieBuilder {
	b.schema.XComponentProps["style"] = style
	return b
}

func (b *chartpieBuilder) AddMeta(key, name string, isDim, isRate bool) *chartpieBuilder {
	meta, ok := b.schema.XComponentProps["meta"].([]map[string]interface{})
	if ok {
		meta = append(meta, map[string]interface{}{
			"id":     key,
			"name":   name,
			"isDim":  isDim,
			"isRate": isRate,
		})
		b.schema.XComponentProps["meta"] = meta
	}
	return b
}
