package field

func Int16(name string) *intBuilder {
	return &intBuilder{&Descriptor{
		Name: name,
		Type: "int16",
	}}
}

func Int32(name string) *intBuilder {
	return &intBuilder{&Descriptor{
		Name: name,
		Type: "int32",
	}}
}

func Int64(name string) *intBuilder {
	return &intBuilder{&Descriptor{
		Name: name,
		Type: "int64",
	}}
}

func BigInt(name string) *intBuilder {
	return &intBuilder{&Descriptor{
		Name: name,
		Type: "bigint",
	}}
}

type intBuilder struct {
	desc *Descriptor
}

func (b *intBuilder) Descriptor() *Descriptor {
	return b.desc
}
