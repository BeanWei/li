package field

func Decimal(name string) *decimalBuilder {
	return &decimalBuilder{&Descriptor{
		Name: name,
		Type: "decimal",
	}}
}

type decimalBuilder struct {
	desc *Descriptor
}

func (b *decimalBuilder) Descriptor() *Descriptor {
	return b.desc
}
