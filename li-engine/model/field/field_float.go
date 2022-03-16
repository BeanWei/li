package field

func Float32(name string) *floatBuilder {
	return &floatBuilder{&Descriptor{
		Name: name,
		Type: "float32",
	}}
}

func Float64(name string) *floatBuilder {
	return &floatBuilder{&Descriptor{
		Name: name,
		Type: "float64",
	}}
}

type floatBuilder struct {
	desc *Descriptor
}

func (b *floatBuilder) Descriptor() *Descriptor {
	return b.desc
}
