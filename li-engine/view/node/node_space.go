package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
	"github.com/gogf/gf/v2/container/gmap"
)

func Space(name string) *spaceBuilder {
	return &spaceBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentSpace,
			XComponentProps: make(map[string]interface{}),
			Properties:      gmap.NewListMap(),
		},
	}}
}

type spaceBuilder struct {
	*NodeBuilder
}

func (b *spaceBuilder) AC(f ac.AC) *spaceBuilder {
	b.schema.AC = f
	return b
}

func (b *spaceBuilder) Title(title string) *spaceBuilder {
	b.schema.Title = title
	return b
}

func (b *spaceBuilder) Description(description string) *spaceBuilder {
	b.schema.Description = description
	return b
}

func (b *spaceBuilder) Align(align string) *spaceBuilder {
	b.schema.XComponentProps["align"] = align
	return b
}

func (b *spaceBuilder) Direction(direction string) *spaceBuilder {
	b.schema.XComponentProps["direction"] = direction
	return b
}

func (b *spaceBuilder) Size(size int) *spaceBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *spaceBuilder) Wrap(wrap bool) *spaceBuilder {
	b.schema.XComponentProps["wrap"] = wrap
	return b
}

func (b *spaceBuilder) SplitByDivider() *spaceBuilder {
	b.schema.XComponentProps["split"] = "divider"
	return b
}
