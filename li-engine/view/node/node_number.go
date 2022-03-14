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

func (b *numberBuilder) Required() *ui.Schema {
	b.schema.Required = true
	return b.schema
}

func (b *numberBuilder) Minimum(min int) *ui.Schema {
	b.schema.Minimum = min
	return b.schema
}

func (b *numberBuilder) Maximum(max int) *ui.Schema {
	b.schema.Maximum = max
	return b.schema
}

func (b *numberBuilder) Title(title string) *ui.Schema {
	b.schema.Title = title
	return b.schema
}

func (b *numberBuilder) Step(step int) *ui.Schema {
	b.schema.XComponentProps["step"] = step
	return b.schema
}

func (b *numberBuilder) Precision(precision int) *ui.Schema {
	b.schema.XComponentProps["precision"] = precision
	return b.schema
}

func (b *numberBuilder) Mode(mode string) *ui.Schema {
	b.schema.XComponentProps["mode"] = mode
	return b.schema
}

func (b *numberBuilder) Prefix(prefix string) *ui.Schema {
	b.schema.XComponentProps["prefix"] = prefix
	return b.schema
}

func (b *numberBuilder) Suffix(suffix string) *ui.Schema {
	b.schema.XComponentProps["suffix"] = suffix
	return b.schema
}
