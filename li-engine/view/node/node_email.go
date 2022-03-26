package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Email(name string) *emailBuilder {
	return &emailBuilder{&textBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeString,
			XComponent:      ui.ComponentInput,
			XComponentProps: make(map[string]interface{}),
			XDecorator:      ui.DecoratorFormItem,
			XValidator:      "email",
		},
	}}}
}

type emailBuilder struct {
	*textBuilder
}
