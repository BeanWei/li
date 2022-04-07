package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ArrayTableMoveUp(name string) *arraytablemoveupBuilder {
	return &arraytablemoveupBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentArrayTableMoveUp,
			XComponentProps: make(map[string]interface{}),
		},
	}}
}

type arraytablemoveupBuilder struct {
	*NodeBuilder
}

func (b *arraytablemoveupBuilder) AC(f ac.AC) *arraytablemoveupBuilder {
	b.schema.AC = f
	return b
}

func (b *arraytablemoveupBuilder) Title(title string) *arraytablemoveupBuilder {
	b.schema.XComponentProps["title"] = title
	return b
}
