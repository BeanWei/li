package node

import (
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/ui"
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
			Properties: make(map[string]*ui.Schema),
		},
	}}
}

type listtableBuilder struct {
	*NodeBuilder
}

func (b *listtableBuilder) LayoutFixed() *listtableBuilder {
	b.schema.XComponentProps["tableLayoutFixed"] = true
	return b
}

func (b *listtableBuilder) Border() *listtableBuilder {
	b.schema.XComponentProps["border"] = true
	return b
}

func (b *listtableBuilder) Hover() *listtableBuilder {
	b.schema.XComponentProps["hover"] = true
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

func (b *listtableBuilder) RowSelectionType(typ string) *listtableBuilder {
	rowsel, ok := b.schema.XComponentProps["rowSelection"].(map[string]interface{})
	if ok {
		rowsel["type"] = typ
		b.schema.XComponentProps["rowSelection"] = rowsel
	}
	return b
}

func (b *listtableBuilder) RowSelectionColumnTitle(title string) *listtableBuilder {
	rowsel, ok := b.schema.XComponentProps["rowSelection"].(map[string]interface{})
	if ok {
		rowsel["columnTitle"] = title
		b.schema.XComponentProps["rowSelection"] = rowsel
	}
	return b
}

func (b *listtableBuilder) RowSelectionColumnWidth(width int) *listtableBuilder {
	rowsel, ok := b.schema.XComponentProps["rowSelection"].(map[string]interface{})
	if ok {
		rowsel["columnWidth"] = width
		b.schema.XComponentProps["rowSelection"] = rowsel
	}
	return b
}

func (b *listtableBuilder) RowSelectionFixed(fixed string) *listtableBuilder {
	rowsel, ok := b.schema.XComponentProps["rowSelection"].(map[string]interface{})
	if ok {
		rowsel["fixed"] = fixed
		b.schema.XComponentProps["rowSelection"] = rowsel
	}
	return b
}

func (b *listtableBuilder) Columns(elements ...view.Node) *listtableBuilder {
	if b.schema.Items == nil {
		b.schema.Items = &ui.Schema{
			Type:       ui.SchemaTypeObject,
			Properties: make(map[string]*ui.Schema),
		}
	}
	for _, element := range elements {
		b.schema.Items.Properties[element.Schema().Name] = element.Schema()
	}
	return b
}
