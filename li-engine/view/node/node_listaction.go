package node

import (
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ListAction(name string) *listactionBuilder {
	return &listactionBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeVoid,
			XComponent: ui.ComponentListAction,
			Properties: make(map[string]*ui.Schema),
		},
	}}
}

type listactionBuilder struct {
	*NodeBuilder
}
