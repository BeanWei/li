package node

import (
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func GridCol(name string) *gridcolBuilder {
	return &gridcolBuilder{schema: &ui.Schema{
		Name:            name,
		Type:            ui.SchemaTypeVoid,
		XComponent:      ui.ComponentGridCol,
		XComponentProps: make(map[string]interface{}),
		Properties:      make(map[string]*ui.Schema),
	}}
}

type gridcolBuilder struct {
	schema *ui.Schema
}

func (b *gridcolBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *gridcolBuilder) Child(elements ...view.Node) *gridcolBuilder {
	for _, element := range elements {
		b.schema.Properties[element.Schema().Name] = element.Schema()
	}
	return b
}

func (b *gridcolBuilder) Content(text string) *gridcolBuilder {
	b.schema.XContent = text
	return b
}

// https://arco.design/react/components/grid#api

// Span 栅格占位格数
func (b *gridcolBuilder) Span(span int) *gridcolBuilder {
	b.schema.XComponentProps["span"] = span
	return b
}

// Offset 栅格左侧的间隔格数，间隔内不可以有栅格
func (b *gridcolBuilder) Offset(offset int) *gridcolBuilder {
	b.schema.XComponentProps["offset"] = offset
	return b
}
