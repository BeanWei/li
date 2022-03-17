package field

import "github.com/BeanWei/li/li-engine/view/node"

func Float32(name string) *floatBuilder {
	return &floatBuilder{&Descriptor{
		Name: name,
		Type: "float32",
		View: node.Number(name).Precision(2),
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
