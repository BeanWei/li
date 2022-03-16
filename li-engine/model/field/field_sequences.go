package field

func Sequences(name string) *sequencesBuilder {
	return &sequencesBuilder{&Descriptor{
		Name: name,
		Type: "sequences",
	}}
}

type sequencesBuilder struct {
	desc *Descriptor
}

func (b *sequencesBuilder) Descriptor() *Descriptor {
	return b.desc
}
