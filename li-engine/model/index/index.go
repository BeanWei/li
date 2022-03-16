package index

type Descriptor struct {
	Fields []string
}

type Builder struct {
	desc *Descriptor
}

func (b *Builder) Descriptor() *Descriptor {
	return b.desc
}

func Fields(fields ...string) *Builder {
	return &Builder{desc: &Descriptor{Fields: fields}}
}
