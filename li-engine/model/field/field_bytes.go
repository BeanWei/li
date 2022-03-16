package field

func Bytes(name string) *bytesBuilder {
	return &bytesBuilder{&Descriptor{
		Name: name,
		Type: "bytes",
	}}
}

type bytesBuilder struct {
	desc *Descriptor
}

func (b *bytesBuilder) Descriptor() *Descriptor {
	return b.desc
}
