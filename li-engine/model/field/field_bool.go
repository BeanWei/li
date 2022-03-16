package field

func Bool(name string) *boolBuilder {
	return &boolBuilder{&Descriptor{
		Name: name,
		Type: "bool",
	}}
}

type boolBuilder struct {
	desc *Descriptor
}

func (b *boolBuilder) Descriptor() *Descriptor {
	return b.desc
}
