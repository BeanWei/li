package field

import "github.com/BeanWei/li/li-engine/view/node"

func Decimal(name string) *decimalBuilder {
	return &decimalBuilder{&Descriptor{
		Name: name,
		Type: "decimal",
		View: node.Number(name).Precision(2).Prefix("¥"),
	}}
}

type decimalBuilder struct {
	desc *Descriptor
}

func (b *decimalBuilder) Descriptor() *Descriptor {
	return b.desc
}
