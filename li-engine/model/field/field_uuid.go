package field

func UUID(name string) *uuidBuilder {
	return &uuidBuilder{&Descriptor{
		Name: name,
		Type: "uuid",
	}}
}

type uuidBuilder struct {
	desc *Descriptor
}

func (b *uuidBuilder) Descriptor() *Descriptor {
	return b.desc
}
