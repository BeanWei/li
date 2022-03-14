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

func (b *percentBuilder) Required() *ui.Schema {
	b.schema.Required = true
	return b.schema
}

func (b *percentBuilder) Minimum(min int) *ui.Schema {
	b.schema.Minimum = min
	return b.schema
}

func (b *percentBuilder) Maximum(max int) *ui.Schema {
	b.schema.Maximum = max
	return b.schema
}

func (b *percentBuilder) Title(title string) *ui.Schema {
	b.schema.Title = title
	return b.schema
}

func (b *percentBuilder) Step(step int) *ui.Schema {
	b.schema.XComponentProps["step"] = step
	return b.schema
}
