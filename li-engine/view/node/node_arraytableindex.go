package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ArrayTableIndex(name string) *arraytableindexBuilder {
	return &arraytableindexBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentArrayTableIndex,
			XComponentProps: make(map[string]interface{}),
		},
	}}
}

type arraytableindexBuilder struct {
	*NodeBuilder
}

func (b *arraytableindexBuilder) AC(f ac.AC) *arraytableindexBuilder {
	b.schema.AC = f
	return b
}

func (b *arraytableindexBuilder) Title(title string) *arraytableindexBuilder {
	b.schema.XComponentProps["title"] = title
	return b
}
