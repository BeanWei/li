package node

import "github.com/BeanWei/li/li-engine/view/ui"

func DropdownMenuURL(name string) *dropdownmenuurlBuilder {
	return &dropdownmenuurlBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentDropdownMenuSubMenu,
			XComponentProps: make(map[string]interface{}),
		},
	}}
}

type dropdownmenuurlBuilder struct {
	*NodeBuilder
}

func (b *dropdownmenuurlBuilder) Title(title string) *dropdownmenuurlBuilder {
	b.schema.Title = title
	return b
}

func (b *dropdownmenuurlBuilder) Href(href string) *dropdownmenuurlBuilder {
	b.schema.XComponentProps["href"] = href
	return b
}
