package node

import (
	"github.com/BeanWei/li/li-engine/view/ui"
	"github.com/gogf/gf/v2/container/gmap"
)

func GridCol(name string) *gridcolBuilder {
	return &gridcolBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			XPath:           name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentGridCol,
			XComponentProps: make(map[string]interface{}),
			Properties:      gmap.NewListMap(),
		},
	}}
}

type gridcolBuilder struct {
	*NodeBuilder
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
