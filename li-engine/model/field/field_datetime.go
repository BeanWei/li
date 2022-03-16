package field

func Datetime(name string) *datetimeBuilder {
	return &datetimeBuilder{&Descriptor{
		Name: name,
		Type: "datetime",
	}}
}

type datetimeBuilder struct {
	desc *Descriptor
}

func (b *datetimeBuilder) Descriptor() *Descriptor {
	return b.desc
}
