package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func Avatar(name string) *avatarBuilder {
	return &avatarBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentAvatar,
			XComponentProps: make(map[string]interface{}),
		},
	}}
}

type avatarBuilder struct {
	*NodeBuilder
}

func (b *avatarBuilder) AC(f ac.AC) *avatarBuilder {
	b.schema.AC = f
	return b
}

func (b *avatarBuilder) Shape(shape string) *avatarBuilder {
	b.schema.XComponentProps["shape"] = shape
	return b
}

func (b *avatarBuilder) Size(size int) *avatarBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *avatarBuilder) Alt(alt string) *avatarBuilder {
	b.schema.XComponentProps["alt"] = alt
	return b
}

func (b *avatarBuilder) Src(src string) *avatarBuilder {
	b.schema.XComponentProps["src"] = src
	return b
}

func (b *avatarBuilder) Style(style map[string]interface{}) *avatarBuilder {
	b.schema.XComponentProps["style"] = style
	return b
}
