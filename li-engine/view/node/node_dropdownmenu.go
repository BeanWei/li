package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/ui"
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

func (b *dropdownmenuBuilder) AC(f ac.AC) *dropdownmenuBuilder {
	b.schema.AC = f
	return b
}

func (b *dropdownmenuBuilder) Droplist(elements ...view.Node) *dropdownmenuBuilder {
	b.Items(elements...)
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
