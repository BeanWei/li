package field

func String(name string) *stringBuilder {
	return &stringBuilder{&Descriptor{
		Name: name,
		Type: "str",
	}}
}

type stringBuilder struct {
	desc *Descriptor
}

func (b *stringBuilder) Descriptor() *Descriptor {
	return b.desc
}
