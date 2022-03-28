package node

import "github.com/BeanWei/li/li-engine/view/ui"

func DropdownMenuSubMenu(name string) *dropdownmenusubmenuBuilder {
	return &dropdownmenusubmenuBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentDropdownMenuSubMenu,
			XComponentProps: make(map[string]interface{}),
		},
	}}
}

type dropdownmenusubmenuBuilder struct {
	*NodeBuilder
}

func (b *dropdownmenusubmenuBuilder) Title(title string) *dropdownmenusubmenuBuilder {
	b.schema.Title = title
	return b
}