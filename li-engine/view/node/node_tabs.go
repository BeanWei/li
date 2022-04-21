package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func Tabs(name string) *tabsBuilder {
	return &tabsBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentTabs,
			XComponentProps: make(map[string]interface{}),
		},
	}}
}

type tabsBuilder struct {
	*NodeBuilder
}

func (b *tabsBuilder) AC(f ac.AC) *tabsBuilder {
	b.schema.AC = f
	return b
}

func (b *tabsBuilder) Title(title string) *tabsBuilder {
	b.SetTitle(title)
	return b
}

func (b *tabsBuilder) DefaultActiveTab(defaultActiveTab string) *tabsBuilder {
	b.schema.XComponentProps["defaultActiveTab"] = defaultActiveTab
	return b
}

func (b *tabsBuilder) Animation(animation bool) *tabsBuilder {
	b.schema.XComponentProps["animation"] = animation
	return b
}

func (b *tabsBuilder) TabPosition(tabPosition string) *tabsBuilder {
	b.schema.XComponentProps["tabPosition"] = tabPosition
	return b
}

func (b *tabsBuilder) Size(size string) *tabsBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *tabsBuilder) Type(typ string) *tabsBuilder {
	b.schema.XComponentProps["type"] = typ
	return b
}

func (b *tabsBuilder) Overflow(overflow string) *tabsBuilder {
	b.schema.XComponentProps["overflow"] = overflow
	return b
}

func (b *tabsBuilder) DestroyOnHide(destroyOnHide bool) *tabsBuilder {
	b.schema.XComponentProps["destroyOnHide"] = destroyOnHide
	return b
}

func (b *tabsBuilder) ScrollPosition(scrollPosition string) *tabsBuilder {
	b.schema.XComponentProps["scrollPosition"] = scrollPosition
	return b
}
