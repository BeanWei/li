package node

import (
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func GridRow(name string) *gridrowBuilder {
	return &gridrowBuilder{schema: &ui.Schema{
		Name:            name,
		Type:            ui.SchemaTypeVoid,
		XComponent:      ui.ComponentGridRow,
		XComponentProps: make(map[string]interface{}),
		Properties:      make(map[string]*ui.Schema),
	}}
}

type gridrowBuilder struct {
	schema *ui.Schema
}

func (b *gridrowBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *gridrowBuilder) Child(elements ...view.Node) *gridrowBuilder {
	for _, element := range elements {
		b.schema.Properties[element.Schema().Name] = element.Schema()
	}
	return b
}

func (b *gridrowBuilder) Content(text string) *gridrowBuilder {
	b.schema.XContent = text
	return b
}

// https://arco.design/react/components/grid#api

// Gutter 栅格间隔
func (b *gridrowBuilder) Gutter(gutter int) *gridrowBuilder {
	b.schema.XComponentProps["gutter"] = gutter
	return b
}

// Align 竖直对齐方式
func (b *gridrowBuilder) Align(align string) *gridrowBuilder {
	b.schema.XComponentProps["align"] = align
	return b
}

// Justify 水平对齐方式
func (b *gridrowBuilder) Justify(justify string) *gridrowBuilder {
	b.schema.XComponentProps["justify"] = justify
	return b
}
