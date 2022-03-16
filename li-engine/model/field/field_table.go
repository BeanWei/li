package field

func Table(name string) *tableBuilder {
	return &tableBuilder{&Descriptor{
		Name: name,
		Type: "array<json>",
	}}
}

type tableBuilder struct {
	desc *Descriptor
}

func (b *tableBuilder) Descriptor() *Descriptor {
	return b.desc
}
