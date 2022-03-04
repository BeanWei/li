package node

func Checkbox(name string) *nodeCheckboxBuilder {
	return &nodeCheckboxBuilder{desc: &Descriptor{
		Name:       name,
		SchemaType: SchemaTypeBool,
		XComponent: "Checkbox",
	}}
}

type nodeCheckboxBuilder struct {
	desc *Descriptor
}

func (b *nodeCheckboxBuilder) Descriptor() *Descriptor {
	return b.desc
}
