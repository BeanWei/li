package field

import "github.com/BeanWei/li/li-engine/view/node"

func Bool(name string) *boolBuilder {
	return &boolBuilder{&Descriptor{
		Name: name,
		Type: "bool",
		View: node.Checkbox(name),
	}}
}

type boolBuilder struct {
	desc *Descriptor
}

func (b *boolBuilder) Descriptor() *Descriptor {
	return b.desc
}
