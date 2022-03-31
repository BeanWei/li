package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ActionForModalCancel(name string) *actionformmodalcancelBuilder {
	return &actionformmodalcancelBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeVoid,
			XComponent: ui.ComponentActionFormModalCancel,
		},
	}}
}

type actionformmodalcancelBuilder struct {
	*NodeBuilder
}

func (b *actionformmodalcancelBuilder) AC(f ac.AC) *actionformmodalcancelBuilder {
	b.schema.AC = f
	return b
}

func (b *actionformmodalcancelBuilder) Title(title string) *actionformmodalcancelBuilder {
	b.schema.Title = title
	return b
}
