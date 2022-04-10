package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Chart(name string) *chartBuilder {
	return &chartBuilder{&formgridBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentChart,
			XComponentProps: make(map[string]interface{}),
		},
	},
	}}
}

type chartBuilder struct {
	*formgridBuilder
}
