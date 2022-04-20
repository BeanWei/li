package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ListActionSearch(name string) *listactionsearchBuilder {
	return &listactionsearchBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeVoid,
			XComponent: ui.ComponentListActionSearch,
		},
	}}
}

type listactionsearchBuilder struct {
	*NodeBuilder
}

func (b *listactionsearchBuilder) AC(f ac.AC) *listactionsearchBuilder {
	b.schema.AC = f
	return b
}

func (b *listactionsearchBuilder) Style(style map[string]interface{}) *listactionsearchBuilder {
	b.schema.XComponentProps["style"] = style
	return b
}

func (b *listactionsearchBuilder) Placeholder(placeholder string) *listactionsearchBuilder {
	b.schema.XComponentProps["placeholder"] = placeholder
	return b
}

func (b *listactionsearchBuilder) Position(position string) *listactionsearchBuilder {
	b.schema.XComponentProps["position"] = position
	return b
}
