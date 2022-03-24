package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Percent(name string) *percentBuilder {
	return &percentBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeNumber,
			XComponent: ui.ComponentInputNumber,
			XComponentProps: map[string]interface{}{
				"suffix":    "%",
				"precision": 2,
				"step":      1,
			},
			XDecorator: ui.DecoratorFormItem,
		},
	}}
}

type percentBuilder struct {
	*NodeBuilder
}

func (b *percentBuilder) Step(step int) *percentBuilder {
	b.schema.XComponentProps["step"] = step
	return b
}
