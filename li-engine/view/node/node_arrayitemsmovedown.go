package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ArrayItemsMoveDown(name string) *arrayitemsmovedownBuilder {
	return &arrayitemsmovedownBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentArrayItemsMoveDown,
			XComponentProps: make(map[string]interface{}),
		},
	}}
}

type arrayitemsmovedownBuilder struct {
	*NodeBuilder
}

func (b *arrayitemsmovedownBuilder) AC(f ac.AC) *arrayitemsmovedownBuilder {
	b.schema.AC = f
	return b
}

func (b *arrayitemsmovedownBuilder) Title(title string) *arrayitemsmovedownBuilder {
	b.schema.XComponentProps["title"] = title
	return b
}
