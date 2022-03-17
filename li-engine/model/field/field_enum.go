package field

import (
	"strings"

	"github.com/BeanWei/li/li-engine/view/node"
	"github.com/gogf/gf/v2/text/gstr"
)

func Enum(name string) *enumBuilder {
	return &enumBuilder{&Descriptor{
		Name:        name,
		Type:        "enum",
		Constraints: make([]string, 0),
		View:        node.Select(name),
	}}
}

type enumBuilder struct {
	desc *Descriptor
}

func (b *enumBuilder) Values(values ...string) *enumBuilder {
	b.desc.Constraints = append(b.desc.Constraints, "one_of('"+strings.Join(values, "', '")+"')")
	enums := make([]map[string]interface{}, len(values))
	for i, v := range values {
		enums[i] = map[string]interface{}{
			v: gstr.CaseCamel(v),
		}
	}
	b.desc.View.Schema().Enum = enums
	return b
}

func (b *enumBuilder) Descriptor() *Descriptor {
	return b.desc
}
