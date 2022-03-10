package field

import (
	"fmt"
	"strings"
)

// Support types: https://www.edgedb.com/docs/datamodel/primitives

// String a variable-length string
func String(name string) *stringBuilder {
	return &stringBuilder{&Descriptor{
		Name: name,
		Type: TypeString,
	}}
}

// Bool logical boolean (true/false)
func Bool(name string) *boolBuilder {
	return &boolBuilder{&Descriptor{
		Name: name,
		Type: TypeBool,
	}}
}

// Int16 16-bit integer
func Int16(name string) *numberBuilder {
	return &numberBuilder{&Descriptor{
		Name: name,
		Type: TypeInt16,
	}}
}

// Int32 32-bit integer
func Int32(name string) *numberBuilder {
	return &numberBuilder{&Descriptor{
		Name: name,
		Type: TypeInt32,
	}}
}

// Int64 64-bit integer
func Int64(name string) *numberBuilder {
	return &numberBuilder{&Descriptor{
		Name: name,
		Type: TypeInt64,
	}}
}

// Float32 32-bit floating point number
func Float32(name string) *numberBuilder {
	return &numberBuilder{&Descriptor{
		Name: name,
		Type: TypeFloat32,
	}}
}

// Float64 64-bit floating point number
func Float64(name string) *numberBuilder {
	return &numberBuilder{&Descriptor{
		Name: name,
		Type: TypeFloat64,
	}}
}

// BigInt arbitrary precision integer
func BigInt(name string) *numberBuilder {
	return &numberBuilder{&Descriptor{
		Name: name,
		Type: TypeBigInt,
	}}
}

// Decimal arbitrary precision number
func Decimal(name string) *numberBuilder {
	return &numberBuilder{&Descriptor{
		Name: name,
		Type: TypeDecimal,
	}}
}

// Map arbitrary JSON data
func Map(name string) *anyBuilder {
	return &anyBuilder{&Descriptor{
		Name: name,
		Type: TypeMap,
	}}
}

// Strings arbitrary JSON data
func Strings(name string) *anyBuilder {
	return &anyBuilder{&Descriptor{
		Name: name,
		Type: TypeStrings,
	}}
}

// Objects arbitrary JSON data
func Objects(name string) *anyBuilder {
	return &anyBuilder{&Descriptor{
		Name: name,
		Type: TypeObjects,
	}}
}

// UUID uuid type
func UUID(name string) *anyBuilder {
	return &anyBuilder{&Descriptor{
		Name: name,
		Type: TypeUUID,
	}}
}

// Bytes arbitrary precision number
func Bytes(name string) *anyBuilder {
	return &anyBuilder{&Descriptor{
		Name: name,
		Type: TypeBytes,
	}}
}

// Datetime Timezone-aware point in time
func Datetime(name string) *anyBuilder {
	return &anyBuilder{&Descriptor{
		Name: name,
		Type: TypeDatetime,
	}}
}

// Duration absolute time span
func Duration(name string) *anyBuilder {
	return &anyBuilder{&Descriptor{
		Name: name,
		Type: TypeDuration,
	}}
}

// Sequences auto-incrementing sequence of int64
func Sequences(name string) *anyBuilder {
	return &anyBuilder{&Descriptor{
		Name: name,
		Type: TypeSequences,
	}}
}

// Enum enum
func Enum(name string) *enumBuilder {
	return &enumBuilder{&Descriptor{
		Name: name,
		Type: TypeEnum,
	}}
}

// Link link field
func Link(name string) *linkBuilder {
	return &linkBuilder{&Descriptor{
		Name: name,
		Type: TypeLink,
		Link: struct {
			IsLink  bool
			IsMulti bool
			From    string
			To      string
		}{
			IsLink: true,
		},
	}}
}

type anyBuilder struct {
	desc *Descriptor
}

func (b *anyBuilder) Descriptor() *Descriptor {
	return b.desc
}

type stringBuilder struct {
	desc *Descriptor
}

func (b *stringBuilder) Required() *stringBuilder {
	b.desc.Required = true
	return b
}

func (b *stringBuilder) Unique() *stringBuilder {
	b.desc.Constraints = append(b.desc.Constraints, "exclusive")
	return b
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

func (b *stringBuilder) Descriptor() *Descriptor {
	return b.desc
}

type boolBuilder struct {
	desc *Descriptor
}

func (b *boolBuilder) Descriptor() *Descriptor {
	return b.desc
}

type numberBuilder struct {
	desc *Descriptor
}

func (b *numberBuilder) Descriptor() *Descriptor {
	return b.desc
}

type enumBuilder struct {
	desc *Descriptor
}

func (b *enumBuilder) Values(values ...string) *enumBuilder {
	b.desc.Constraints = append(b.desc.Constraints, "one_of('"+strings.Join(values, "', '")+"')")
	return b
}

func (b *enumBuilder) Descriptor() *Descriptor {
	return b.desc
}

type linkBuilder struct {
	desc *Descriptor
}

func (b *linkBuilder) Multi(t string) *Descriptor {
	b.desc.Link.To = t
	return b.desc
}

func (b *linkBuilder) From(typ string, field string) *Descriptor {
	b.desc.Link.To = ".<" + field + "[is" + typ + "]"
	return b.desc
}

func (b *linkBuilder) To(t string) *Descriptor {
	b.desc.Link.To = t
	return b.desc
}

func (b *linkBuilder) Descriptor() *Descriptor {
	return b.desc
}

// A Descriptor for field configuration.
type Descriptor struct {
	Name      string
	Type      Type
	Sensitive bool
	Required  bool
	Default   interface{}
	Link      struct {
		IsLink  bool
		IsMulti bool
		From    string
		To      string
	}
	Constraints []string
}

/*
type Descriptor struct {
	// 公共属性
	Name     string
	Names    [2]string
	Required bool
	Multi    bool
	Default  interface{}
	Defaults [2]string

	// 实体属性
	ValueType   Type
	EdgeName    string
	Sensitive   bool
	Constraints []string

	// 视图属性
	ViewType        string
	ViewUIProps     map[string]interface{}
	ViewSchemaProps map[string]interface{}
}
func (d *Descriptor) ToESDL() string {
	var b strings.Builder
	if d.ValueType == TypeLink || d.ValueType == TypeMultiLink {
		b.WriteString(d.ValueType.String())
	} else {
		b.WriteString("property ")
	}
	if d.EdgeName != "" {
		b.WriteString(" -> " + d.EdgeName)
	} else {
		b.WriteString(" -> " + d.ValueType.String())
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
*/

func (d *Descriptor) ToESDL() string {
	var b strings.Builder
	if d.Link.IsLink {
		if d.Link.IsMulti {
			b.WriteString("multi ")
		}
		b.WriteString("link ")
	} else {
		b.WriteString("property ")
	}
	b.WriteString(d.Name)
	if d.Link.IsLink {
		if d.Link.To != "" {
			b.WriteString(" -> " + d.Link.To)
		} else if d.Link.From != "" {
			b.WriteString(" := " + d.Link.From)
		}
	} else {
		b.WriteString(" -> " + d.Type.String())
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
