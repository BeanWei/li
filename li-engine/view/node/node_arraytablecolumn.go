package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/ui"
	"github.com/gogf/gf/v2/container/gmap"
)

func ArrayTableColumn(name string) *arraytablecolumnBuilder {
	return &arraytablecolumnBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentArrayTableColumn,
			XComponentProps: make(map[string]interface{}),
			Properties:      gmap.NewListMap(),
		},
	}}
}

type arraytablecolumnBuilder struct {
	*NodeBuilder
}

func (b *arraytablecolumnBuilder) AC(f ac.AC) *arraytablecolumnBuilder {
	b.schema.AC = f
	return b
}

func (b *arraytablecolumnBuilder) Title(title string) *arraytablecolumnBuilder {
	b.schema.XComponentProps["title"] = title
	return b
}

func (b *arraytablecolumnBuilder) DataIndex(dataIndex string) *arraytablecolumnBuilder {
	b.schema.XComponentProps["dataIndex"] = dataIndex
	return b
}

func (b *arraytablecolumnBuilder) Align(align string) *arraytablecolumnBuilder {
	b.schema.XComponentProps["align"] = align
	return b
}

func (b *arraytablecolumnBuilder) Ellipsis(ellipsis bool) *arraytablecolumnBuilder {
	b.schema.XComponentProps["ellipsis"] = ellipsis
	return b
}

func (b *arraytablecolumnBuilder) Width(width int) *arraytablecolumnBuilder {
	b.schema.XComponentProps["width"] = width
	return b
}

func (b *arraytablecolumnBuilder) Render(element view.Node) *arraytablecolumnBuilder {
	b.schema.Properties.Set(element.Schema().Name, element.Schema())
	return b
}
