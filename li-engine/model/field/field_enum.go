package field

func Enum(name string) *enumBuilder {
	return &enumBuilder{&Descriptor{
		Name: name,
		Type: "enum",
	}}
}

type enumBuilder struct {
	desc *Descriptor
}

func (b *enumBuilder) Descriptor() *Descriptor {
	return b.desc
}
