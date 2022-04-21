package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ActionFormDrawerCancel(name string) *actionformdrawercancelBuilder {
	return &actionformdrawercancelBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeVoid,
			XComponent: ui.ComponentActionFormDrawerCancel,
		},
	}}
}

type actionformdrawercancelBuilder struct {
	*NodeBuilder
}

func (b *actionformdrawercancelBuilder) AC(f ac.AC) *actionformdrawercancelBuilder {
	b.schema.AC = f
	return b
}

func (b *actionformdrawercancelBuilder) Title(title string) *actionformdrawercancelBuilder {
	b.SetTitle(title)
	return b
}
