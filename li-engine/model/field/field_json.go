package field

func JSON(name string) *jsonBuilder {
	return &jsonBuilder{&Descriptor{
		Name: name,
		Type: "json",
	}}
}

type jsonBuilder struct {
	desc *Descriptor
}

func (b *jsonBuilder) Descriptor() *Descriptor {
	return b.desc
}
