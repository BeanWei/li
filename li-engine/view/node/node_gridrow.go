package node

import (
	"github.com/BeanWei/li/li-engine/view/ui"
	"github.com/gogf/gf/v2/container/gmap"
)

func GridRow(name string) *gridrowBuilder {
	return &gridrowBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			XPath:           name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentGridRow,
			XComponentProps: make(map[string]interface{}),
			Properties:      gmap.NewListMap(),
		},
	}}
}

type gridrowBuilder struct {
	*NodeBuilder
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
