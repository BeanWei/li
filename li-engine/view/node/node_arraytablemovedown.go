package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ArrayTableMoveDown(name string) *arraytablemovedownBuilder {
	return &arraytablemovedownBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentArrayTableMoveDown,
			XComponentProps: make(map[string]interface{}),
		},
	}}
}

type arraytablemovedownBuilder struct {
	*NodeBuilder
}

func (b *arraytablemovedownBuilder) AC(f ac.AC) *arraytablemovedownBuilder {
	b.schema.AC = f
	return b
}

func (b *arraytablemovedownBuilder) Title(title string) *arraytablemovedownBuilder {
	b.schema.XComponentProps["title"] = title
	return b
}
