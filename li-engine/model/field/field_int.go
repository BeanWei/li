package field

import "github.com/BeanWei/li/li-engine/view/node"

func Int16(name string) *intBuilder {
	return &intBuilder{&Descriptor{
		Name: name,
		Type: "int16",
		View: node.Number(name).Precision(2),
	}}
}

func Int32(name string) *intBuilder {
	return &intBuilder{&Descriptor{
		Name: name,
		Type: "int32",
		View: node.Number(name).Precision(2),
	}}
}

func Int64(name string) *intBuilder {
	return &intBuilder{&Descriptor{
		Name: name,
		Type: "int64",
		View: node.Number(name).Precision(2),
	}}
}

func BigInt(name string) *intBuilder {
	return &intBuilder{&Descriptor{
		Name: name,
		Type: "bigint",
		View: node.Number(name).Precision(2),
	}}
}

type intBuilder struct {
	desc *Descriptor
}

func (b *intBuilder) Descriptor() *Descriptor {
	return b.desc
}
