package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/controller"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ChartLine(name string) *chartlineBuilder {
	return &chartlineBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeVoid,
			XComponent: ui.ComponentChartLine,
			XComponentProps: map[string]interface{}{
				"meta": make([]map[string]interface{}, 0),
			},
			XDecorator:      ui.DecoratorChartItem,
			XDecoratorProps: make(map[string]interface{}),
		},
	}}
}

type chartlineBuilder struct {
	*NodeBuilder
}

func (b *chartlineBuilder) AC(f ac.AC) *chartlineBuilder {
	b.schema.AC = f
	return b
}

func (b *chartlineBuilder) Title(title string) *chartlineBuilder {
	b.schema.XDecoratorProps["title"] = title
	return b
}

func (b *chartlineBuilder) SubTitle(subTitle string) *chartlineBuilder {
	b.schema.XDecoratorProps["subTitle"] = subTitle
	return b
}

func (b *chartlineBuilder) Description(description string) *chartlineBuilder {
	b.schema.Description = description
	return b
}

func (b *chartlineBuilder) ForInit(operation string, handler interface{}) *chartlineBuilder {
	b.schema.XDecoratorProps["forInit"] = operation
	b.schema.HandlerNames = append(b.schema.HandlerNames, operation)
	controller.Bind(operation, handler)
	return b
}

func (b *chartlineBuilder) ForInitVariables(variables map[string]interface{}) *chartlineBuilder {
	b.schema.XDecoratorProps["forInitVariables"] = variables
	return b
}

func (b *chartlineBuilder) GridSpan(gridSpan int) *chartlineBuilder {
	b.schema.XDecoratorProps["gridSpan"] = gridSpan
	return b
}

func (b *chartlineBuilder) Style(style map[string]interface{}) *chartlineBuilder {
	b.schema.XComponentProps["style"] = style
	return b
}

func (b *chartlineBuilder) AddMeta(key, name string, isDim, isRate bool) *chartlineBuilder {
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

func (b *chartlineBuilder) WithArea(withArea bool) *chartlineBuilder {
	b.schema.XComponentProps["withArea"] = withArea
	return b
}
