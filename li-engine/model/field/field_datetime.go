package field

import "github.com/BeanWei/li/li-engine/view/node"

func Datetime(name string) *datetimeBuilder {
	return &datetimeBuilder{&Descriptor{
		Name: name,
		Type: "datetime",
		View: node.Date(name).Mode("time"),
	}}
}

type datetimeBuilder struct {
	desc *Descriptor
}

func (b *datetimeBuilder) Descriptor() *Descriptor {
	return b.desc
}
