package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
	"github.com/gogf/gf/v2/container/gmap"
)

func ListAction(name string) *listactionBuilder {
	return &listactionBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeVoid,
			XComponent: ui.ComponentListAction,
			Properties: gmap.NewListMap(),
		},
	}}
}

type listactionBuilder struct {
	*NodeBuilder
}

func (b *listactionBuilder) AC(f ac.AC) *listactionBuilder {
	b.schema.AC = f
	return b
}

func (b *listactionBuilder) Title(title string) *listactionBuilder {
	b.SetTitle(title)
	return b
}

func (b *listactionBuilder) Description(description string) *listactionBuilder {
	b.SetDescription(description)
	return b
}
