package node

import (
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/ui"
	"github.com/gogf/gf/v2/container/gmap"
)

func DropdownMenu(name string) *dropdownmenuBuilder {
	return &dropdownmenuBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeVoid,
			XComponent: ui.ComponentDropdownMenu,
			XComponentProps: map[string]interface{}{
				"triggerProps": map[string]interface{}{
					"popupStyle": make(map[string]interface{}),
				},
			},
		},
	}}
}

type dropdownmenuBuilder struct {
	*NodeBuilder
}

func (b *dropdownmenuBuilder) Droplist(elements ...view.Node) *dropdownmenuBuilder {
	if b.schema.Items == nil {
		b.schema.Items = &ui.Schema{
			Type:       ui.SchemaTypeVoid,
			Properties: gmap.NewListMap(),
		}
	}
	for _, element := range elements {
		b.schema.Items.Properties.Set(element.Schema().Name, element.Schema())
	}
	return b
}

func (b *dropdownmenuBuilder) Position(position string) *dropdownmenuBuilder {
	b.schema.XComponentProps["position"] = position
	return b
}

func (b *dropdownmenuBuilder) Trigger(trigger string) *dropdownmenuBuilder {
	b.schema.XComponentProps["trigger"] = trigger
	return b
}

func (b *dropdownmenuBuilder) TriggerPopupStyle(style map[string]interface{}) *dropdownmenuBuilder {
	triggerProps, ok := b.schema.XComponentProps["triggerProps"].(map[string]interface{})
	if !ok {
		return b
	}
	triggerProps["popupStyle"] = style
	b.schema.XComponentProps["triggerProps"] = triggerProps
	return b
}
