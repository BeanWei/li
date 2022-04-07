package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ArrayItemsMoveUp(name string) *arrayitemsmoveupBuilder {
	return &arrayitemsmoveupBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentArrayItemsMoveUp,
			XComponentProps: make(map[string]interface{}),
		},
	}}
}

type arrayitemsmoveupBuilder struct {
	*NodeBuilder
}

func (b *arrayitemsmoveupBuilder) AC(f ac.AC) *arrayitemsmoveupBuilder {
	b.schema.AC = f
	return b
}

func (b *arrayitemsmoveupBuilder) Title(title string) *arrayitemsmoveupBuilder {
	b.schema.XComponentProps["title"] = title
	return b
}
