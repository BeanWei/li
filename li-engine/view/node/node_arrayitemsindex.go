package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ArrayItemsIndex(name string) *arrayitemsindexBuilder {
	return &arrayitemsindexBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentArrayItemsIndex,
			XComponentProps: make(map[string]interface{}),
		},
	}}
}

type arrayitemsindexBuilder struct {
	*NodeBuilder
}

func (b *arrayitemsindexBuilder) AC(f ac.AC) *arrayitemsindexBuilder {
	b.schema.AC = f
	return b
}

func (b *arrayitemsindexBuilder) Title(title string) *arrayitemsindexBuilder {
	b.schema.XComponentProps["title"] = title
	return b
}
