package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ArrayItemsSortHandle(name string) *arrayitemssorthandleBuilder {
	return &arrayitemssorthandleBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentArrayItemsSortHandle,
			XComponentProps: make(map[string]interface{}),
		},
	}}
}

type arrayitemssorthandleBuilder struct {
	*NodeBuilder
}

func (b *arrayitemssorthandleBuilder) AC(f ac.AC) *arrayitemssorthandleBuilder {
	b.schema.AC = f
	return b
}

func (b *arrayitemssorthandleBuilder) Title(title string) *arrayitemssorthandleBuilder {
	b.schema.XComponentProps["title"] = title
	return b
}
