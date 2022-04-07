package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/ui"
	"github.com/gogf/gf/v2/container/gmap"
)

func ArrayTable(name string) *arraytableBuilder {
	return &arraytableBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeArray,
			XComponent:      ui.ComponentArrayTable,
			XComponentProps: make(map[string]interface{}),
			XDecorator:      ui.DecoratorFormItem,
			Properties:      gmap.NewListMap(),
		},
	}}
}

type arraytableBuilder struct {
	*NodeBuilder
}

func (b *arraytableBuilder) AC(f ac.AC) *arraytableBuilder {
	b.schema.AC = f
	return b
}

func (b *arraytableBuilder) LayoutFixed(tableLayoutFixed bool) *arraytableBuilder {
	b.schema.XComponentProps["tableLayoutFixed"] = tableLayoutFixed
	return b
}

func (b *arraytableBuilder) Border(border bool) *arraytableBuilder {
	b.schema.XComponentProps["border"] = border
	return b
}

func (b *arraytableBuilder) Hover(hover bool) *arraytableBuilder {
	b.schema.XComponentProps["hover"] = hover
	return b
}

func (b *arraytableBuilder) Stripe(stripe bool) *arraytableBuilder {
	b.schema.XComponentProps["stripe"] = stripe
	return b
}

func (b *arraytableBuilder) Size(size string) *arraytableBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *arraytableBuilder) Pagination(pagination interface{}) *arraytableBuilder {
	b.schema.XComponentProps["pagination"] = pagination
	return b
}

func (b *arraytableBuilder) Scroll(scroll map[string]interface{}) *arraytableBuilder {
	b.schema.XComponentProps["scroll"] = scroll
	return b
}

func (b *arraytableBuilder) Columns(elements ...view.Node) *arraytableBuilder {
	b.Items(elements...)
	return b
}
