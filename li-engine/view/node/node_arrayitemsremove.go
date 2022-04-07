package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ArrayItemRemove(name string) *arrayitemsremoveBuilder {
	return &arrayitemsremoveBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentArrayItemsRemove,
			XComponentProps: make(map[string]interface{}),
		},
	}}
}

type arrayitemsremoveBuilder struct {
	*NodeBuilder
}

func (b *arrayitemsremoveBuilder) AC(f ac.AC) *arrayitemsremoveBuilder {
	b.schema.AC = f
	return b
}

func (b *arrayitemsremoveBuilder) Title(title string) *arrayitemsremoveBuilder {
	b.schema.XComponentProps["title"] = title
	return b
}
