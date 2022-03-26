package node

import "github.com/BeanWei/li/li-engine/view/ui"

func ListCard(name string) *listcardBuilder {
	return &listcardBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentListCard,
			XComponentProps: make(map[string]interface{}),
			Properties:      make(map[string]*ui.Schema),
		},
	}}
}

type listcardBuilder struct {
	*NodeBuilder
}

func (b *listcardBuilder) Title(title string) *listcardBuilder {
	b.schema.Title = title
	return b
}

func (b *listcardBuilder) Description(description string) *listcardBuilder {
	b.schema.Description = description
	return b
}
