package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Number(name string) *numberBuilder {
	return &numberBuilder{schema: &ui.Schema{
		Name:            name,
		Type:            ui.SchemaTypeNumber,
		XComponent:      ui.ComponentInputNumber,
		XComponentProps: make(map[string]interface{}),
		XDecorator:      ui.DecoratorFormItem,
	}}
}

type numberBuilder struct {
	schema *ui.Schema
}

func (b *numberBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *numberBuilder) Required() *numberBuilder {
	b.schema.Required = true
	return b
}

func (b *numberBuilder) Minimum(min int) *numberBuilder {
	b.schema.Minimum = min
	return b
}

func (b *numberBuilder) Maximum(max int) *numberBuilder {
	b.schema.Maximum = max
	return b
}

func (b *numberBuilder) Title(title string) *numberBuilder {
	b.schema.Title = title
	return b
}

func (b *numberBuilder) Step(step int) *numberBuilder {
	b.schema.XComponentProps["step"] = step
	return b
}

func (b *numberBuilder) Precision(precision int) *numberBuilder {
	b.schema.XComponentProps["precision"] = precision
	return b
}

func (b *numberBuilder) Mode(mode string) *numberBuilder {
	b.schema.XComponentProps["mode"] = mode
	return b
}

func (b *numberBuilder) Prefix(prefix string) *numberBuilder {
	b.schema.XComponentProps["prefix"] = prefix
	return b
}

func (b *numberBuilder) Suffix(suffix string) *numberBuilder {
	b.schema.XComponentProps["suffix"] = suffix
	return b
}
