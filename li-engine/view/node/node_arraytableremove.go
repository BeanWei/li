package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ArrayTableRemove(name string) *arraytableremoveBuilder {
	return &arraytableremoveBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentArrayTableRemove,
			XComponentProps: make(map[string]interface{}),
		},
	}}
}

type arraytableremoveBuilder struct {
	*NodeBuilder
}

func (b *arraytableremoveBuilder) AC(f ac.AC) *arraytableremoveBuilder {
	b.schema.AC = f
	return b
}

func (b *arraytableremoveBuilder) Title(title string) *arraytableremoveBuilder {
	b.schema.XComponentProps["title"] = title
	return b
}
