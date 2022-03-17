package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Percent(name string) *percentBuilder {
	return &percentBuilder{schema: &ui.Schema{
		Name:       name,
		Type:       ui.SchemaTypeNumber,
		XComponent: ui.ComponentInputNumber,
		XComponentProps: map[string]interface{}{
			"suffix":    "%",
			"precision": 2,
			"step":      1,
		},
		XDecorator: ui.DecoratorFormItem,
	}}
}

type percentBuilder struct {
	schema *ui.Schema
}

func (b *percentBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *percentBuilder) Required() *percentBuilder {
	b.schema.Required = true
	return b
}

func (b *percentBuilder) Minimum(min int) *percentBuilder {
	b.schema.Minimum = min
	return b
}

func (b *percentBuilder) Maximum(max int) *percentBuilder {
	b.schema.Maximum = max
	return b
}

func (b *percentBuilder) Title(title string) *percentBuilder {
	b.schema.Title = title
	return b
}

func (b *percentBuilder) Step(step int) *percentBuilder {
	b.schema.XComponentProps["step"] = step
	return b
}
