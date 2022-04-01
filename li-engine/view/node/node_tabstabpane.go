package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func TabsTabPane(name string) *tabstabpaneBuilder {
	return &tabstabpaneBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentTabsTabPane,
			XComponentProps: make(map[string]interface{}),
		},
	}}
}

type tabstabpaneBuilder struct {
	*NodeBuilder
}

func (b *tabstabpaneBuilder) AC(f ac.AC) *tabstabpaneBuilder {
	b.schema.AC = f
	return b
}

func (b *tabstabpaneBuilder) Title(title string) *tabstabpaneBuilder {
	b.schema.Title = title
	return b
}

func (b *tabsBuilder) Icon(icon string) *tabsBuilder {
	b.schema.XComponentProps["icon"] = icon
	return b
}
