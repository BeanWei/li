package field

func Strings(name string) *stringsBuilder {
	return &stringsBuilder{&Descriptor{
		Name: name,
		Type: "array<str>",
	}}
}

type stringsBuilder struct {
	desc *Descriptor
}

func (b *stringsBuilder) Descriptor() *Descriptor {
	return b.desc
}
