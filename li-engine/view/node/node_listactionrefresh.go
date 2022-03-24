package node

import "github.com/BeanWei/li/li-engine/view/ui"

func ListActionRefresh(name string) *listactionrefreshBuilder {
	return &listactionrefreshBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeVoid,
			XComponent: ui.ComponentListActionRefresh,
		},
	}}
}

type listactionrefreshBuilder struct {
	*NodeBuilder
}
