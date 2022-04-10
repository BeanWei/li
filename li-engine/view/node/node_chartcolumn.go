package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/controller"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ChartColumn(name string) *chartcolumnBuilder {
	return &chartcolumnBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeVoid,
			XComponent: ui.ComponentChartColumn,
			XComponentProps: map[string]interface{}{
				"meta": make([]map[string]interface{}, 0),
			},
			XDecorator:      ui.DecoratorChartItem,
			XDecoratorProps: make(map[string]interface{}),
		},
	}}
}

type chartcolumnBuilder struct {
	*NodeBuilder
}

func (b *chartcolumnBuilder) AC(f ac.AC) *chartcolumnBuilder {
	b.schema.AC = f
	return b
}

func (b *chartcolumnBuilder) Title(title string) *chartcolumnBuilder {
	b.schema.XDecoratorProps["title"] = title
	return b
}

func (b *chartcolumnBuilder) SubTitle(subTitle string) *chartcolumnBuilder {
	b.schema.XDecoratorProps["subTitle"] = subTitle
	return b
}

func (b *chartcolumnBuilder) Description(description string) *chartcolumnBuilder {
	b.schema.Description = description
	return b
}

func (b *chartcolumnBuilder) ForInit(operation string, handler interface{}) *chartcolumnBuilder {
	b.schema.XDecoratorProps["forInit"] = operation
	b.schema.HandlerNames = append(b.schema.HandlerNames, operation)
	controller.Bind(operation, handler)
	return b
}

func (b *chartcolumnBuilder) ForInitVariables(variables map[string]interface{}) *chartcolumnBuilder {
	b.schema.XDecoratorProps["forInitVariables"] = variables
	return b
}

func (b *chartcolumnBuilder) GridSpan(gridSpan int) *chartcolumnBuilder {
	b.schema.XDecoratorProps["gridSpan"] = gridSpan
	return b
}

func (b *chartcolumnBuilder) Style(style map[string]interface{}) *chartcolumnBuilder {
	b.schema.XComponentProps["style"] = style
	return b
}

func (b *chartcolumnBuilder) AddMeta(key, name string, isDim, isRate bool) *chartcolumnBuilder {
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

func (b *chartcolumnBuilder) Inverted(inverted bool) *chartcolumnBuilder {
	b.schema.XComponentProps["inverted"] = inverted
	return b
}
