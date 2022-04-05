package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/ui"
	"github.com/gogf/gf/v2/container/gmap"
)

func ListTable(name string) *listtableBuilder {
	return &listtableBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeArray,
			XComponent: ui.ComponentListTable,
			XComponentProps: map[string]interface{}{
				"rowSelection": make(map[string]interface{}),
			},
			Properties: gmap.NewListMap(),
		},
	}}
}

type listtableBuilder struct {
	*NodeBuilder
}

func (b *listtableBuilder) AC(f ac.AC) *listtableBuilder {
	b.schema.AC = f
	return b
}

func (b *listtableBuilder) LayoutFixed(tableLayoutFixed bool) *listtableBuilder {
	b.schema.XComponentProps["tableLayoutFixed"] = tableLayoutFixed
	return b
}

func (b *listtableBuilder) Border(border bool) *listtableBuilder {
	b.schema.XComponentProps["border"] = border
	return b
}

func (b *listtableBuilder) Hover(hover bool) *listtableBuilder {
	b.schema.XComponentProps["hover"] = hover
	return b
}

func (b *listtableBuilder) Stripe(stripe bool) *listtableBuilder {
	b.schema.XComponentProps["stripe"] = stripe
	return b
}

func (b *listtableBuilder) Size(size string) *listtableBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *listtableBuilder) ActionBar(element view.Node) *listtableBuilder {
	b.schema.Properties.Set(element.Schema().Name, element.Schema())
	return b
}

func (b *listtableBuilder) Columns(elements ...view.Node) *listtableBuilder {
	b.Items(elements...)
	return b
}
