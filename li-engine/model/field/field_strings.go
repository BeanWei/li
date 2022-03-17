package field

import "github.com/BeanWei/li/li-engine/view/node"

func Strings(name string) *stringsBuilder {
	return &stringsBuilder{&Descriptor{
		Name: name,
		Type: "array<str>",
		View: node.Tags(name),
	}}
}

type stringsBuilder struct {
	desc *Descriptor
}

func (b *stringsBuilder) Descriptor() *Descriptor {
	return b.desc
}
