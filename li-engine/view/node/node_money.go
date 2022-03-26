package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Money(name string) *moneyBuilder {
	return &moneyBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeString,
			XComponent:      ui.ComponentMoney,
			XComponentProps: make(map[string]interface{}),
			XDecorator:      ui.DecoratorFormItem,
		},
	}}
}

type moneyBuilder struct {
	*NodeBuilder
}

func (b *moneyBuilder) Title(title string) *moneyBuilder {
	b.schema.Title = title
	return b
}

func (b *moneyBuilder) Description(description string) *moneyBuilder {
	b.schema.Description = description
	return b
}

func (b *moneyBuilder) Default(value interface{}) *moneyBuilder {
	b.schema.Default = value
	return b
}

func (b *moneyBuilder) Step(step int) *moneyBuilder {
	b.schema.XComponentProps["step"] = step
	return b
}

func (b *moneyBuilder) Precision(precision int) *moneyBuilder {
	b.schema.XComponentProps["precision"] = precision
	return b
}

func (b *moneyBuilder) Currency(currency string) *moneyBuilder {
	b.schema.XComponentProps["currency"] = currency
	return b
}
