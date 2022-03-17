package field

import (
	"fmt"

	"github.com/BeanWei/li/li-engine/view/node"
)

func String(name string) *stringBuilder {
	return &stringBuilder{&Descriptor{
		Name:        name,
		Type:        "str",
		Constraints: make([]string, 0),
		View:        node.Text(name),
	}}
}

func LongString(name string) *longstringBuilder {
	return &longstringBuilder{&Descriptor{
		Name: name,
		Type: "str",
		View: node.TextArea(name),
	}}
}

type stringBuilder struct {
	desc *Descriptor
}

func (b *stringBuilder) Descriptor() *Descriptor {
	return b.desc
}

func (b *stringBuilder) Required() *Descriptor {
	b.desc.Required = true
	return b.desc
}

func (b *stringBuilder) Unique() *Descriptor {
	b.desc.Constraints = append(b.desc.Constraints, "exclusive")
	return b.desc
}

func (b *stringBuilder) Sensitive() *stringBuilder {
	b.desc.Sensitive = true
	return b
}

func (b *stringBuilder) Match(re string) *stringBuilder {
	b.desc.Constraints = append(b.desc.Constraints, `regexp(r"`+re+")")
	return b
}

func (b *stringBuilder) MaxLen(i int) *stringBuilder {
	b.desc.Constraints = append(b.desc.Constraints, fmt.Sprintf("max_len_value(%d)", i))
	return b
}

func (b *stringBuilder) MinLen(i int) *stringBuilder {
	b.desc.Constraints = append(b.desc.Constraints, fmt.Sprintf("min_len_value(%d)", i))
	return b
}

func (b *stringBuilder) ExpressionOn(expr string) *stringBuilder {
	b.desc.Constraints = append(b.desc.Constraints, "expression on ("+expr+")")
	return b
}

func (b *stringBuilder) Default(s string) *stringBuilder {
	b.desc.Default = s
	return b
}

type longstringBuilder struct {
	desc *Descriptor
}

func (b *longstringBuilder) Descriptor() *Descriptor {
	return b.desc
}

func (b *longstringBuilder) Required() *Descriptor {
	b.desc.Required = true
	return b.desc
}

func (b *longstringBuilder) MaxLen(i int) *longstringBuilder {
	b.desc.Constraints = append(b.desc.Constraints, fmt.Sprintf("max_len_value(%d)", i))
	return b
}

func (b *longstringBuilder) MinLen(i int) *longstringBuilder {
	b.desc.Constraints = append(b.desc.Constraints, fmt.Sprintf("min_len_value(%d)", i))
	return b
}
