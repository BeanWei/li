package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ArrayTableSortHandle(name string) *arraytablesorthandleBuilder {
	return &arraytablesorthandleBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentArrayTableSortHandle,
			XComponentProps: make(map[string]interface{}),
		},
	}}
}

type arraytablesorthandleBuilder struct {
	*NodeBuilder
}

func (b *arraytablesorthandleBuilder) AC(f ac.AC) *arraytablesorthandleBuilder {
	b.schema.AC = f
	return b
}

func (b *arraytablesorthandleBuilder) Title(title string) *arraytablesorthandleBuilder {
	b.schema.XComponentProps["title"] = title
	return b
}
