package node

import (
	"github.com/BeanWei/li/li-engine/view/ui"
	"github.com/gogf/gf/v2/container/gmap"
)

func ListAction(name string) *listactionBuilder {
	return &listactionBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			XPath:      name,
			Type:       ui.SchemaTypeVoid,
			XComponent: ui.ComponentListAction,
			Properties: gmap.NewListMap(),
		},
	}}
}

type listactionBuilder struct {
	*NodeBuilder
}

func (b *listactionBuilder) Title(title string) *listactionBuilder {
	b.schema.Title = title
	return b
}

func (b *listactionBuilder) Description(description string) *listactionBuilder {
	b.schema.Description = description
	return b
}
