package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Divider(name string) *dividerBuilder {
	return &dividerBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			XPath:           name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentDivider,
			XComponentProps: make(map[string]interface{}),
		},
	}}
}

type dividerBuilder struct {
	*NodeBuilder
}

func (b *dividerBuilder) Type(typ string) *dividerBuilder {
	b.schema.XComponentProps["type"] = typ
	return b
}

func (b *dividerBuilder) Orientation(orientation string) *dividerBuilder {
	b.schema.XComponentProps["orientation"] = orientation
	return b
}

func (b *dividerBuilder) Style(style map[string]interface{}) *dividerBuilder {
	b.schema.XComponentProps["style"] = style
	return b
}
