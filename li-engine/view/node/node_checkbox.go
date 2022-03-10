package node

func Checkbox(name string) *nodeCheckboxBuilder {
	return &nodeCheckboxBuilder{desc: &Descriptor{
		Name:          name,
		SchemaType:    SchemaTypeBool,
		ComponentName: "Checkbox",
	}}
}

type nodeCheckboxBuilder struct {
	desc *Descriptor
}

func (b *nodeCheckboxBuilder) SchemaProps(props *SchemaProps) *nodeCheckboxBuilder {
	b.desc.SchemaProps = props
	return b
}

func (b *nodeCheckboxBuilder) Descriptor() *Descriptor {
	return b.desc
}
