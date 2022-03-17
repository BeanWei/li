package edge

import (
	"reflect"
	"strings"
)

func To(name string, t interface{}) *assocBuilder {
	return &assocBuilder{desc: &Descriptor{
		Name: name,
		Type: typ(t),
	}}
}

func From(name string, t interface{}) *inverseBuilder {
	return &inverseBuilder{desc: &Descriptor{
		Name:    name,
		Type:    typ(t),
		Inverse: true,
	}}
}

type Descriptor struct {
	Name        string
	Type        string
	Multi       bool
	Inverse     bool
	Required    bool
	Constraints []string
}

type assocBuilder struct {
	desc *Descriptor
}

func (b *assocBuilder) Descriptor() *Descriptor {
	return b.desc
}

func (b *assocBuilder) Multi() *assocBuilder {
	b.desc.Multi = true
	return b
}

func (b *assocBuilder) Required() *assocBuilder {
	b.desc.Required = true
	return b
}

type inverseBuilder struct {
	desc *Descriptor
}

func (b *inverseBuilder) Descriptor() *Descriptor {
	return b.desc
}

func (b *inverseBuilder) Multi() *inverseBuilder {
	b.desc.Multi = true
	return b
}

func (b *inverseBuilder) Required() *inverseBuilder {
	b.desc.Required = true
	return b
}

func typ(t interface{}) string {
	if rt := reflect.TypeOf(t); rt.NumIn() > 0 {
		return rt.In(0).Name()
	}
	return ""
}

func (d *Descriptor) ToESDL() string {
	var b strings.Builder
	if d.Multi {
		b.WriteString("multi ")
	}
	b.WriteString("link ")
	if d.Inverse {
		b.WriteString(" := " + d.Type)
	} else {
		b.WriteString(" -> " + d.Type)
	}
	if len(d.Constraints) != 0 {
		b.WriteString(" {")
		for _, cst := range d.Constraints {
			b.WriteString(" constraint " + cst + ";")
		}
		b.WriteString(" }")
	} else {
		b.WriteString(";")
	}
	return b.String()
}
