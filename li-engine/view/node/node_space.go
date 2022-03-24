package node

import (
	"github.com/BeanWei/li/li-engine/view/ui"
)

func Space(name string) *spaceBuilder {
	return &spaceBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeVoid,
			XComponent: ui.ComponentSpace,
			Properties: make(map[string]*ui.Schema),
		},
	}}
}

type spaceBuilder struct {
	*NodeBuilder
}
