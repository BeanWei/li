package field

import (
	"strings"

	"github.com/BeanWei/li/li-engine/view"
)

type Descriptor struct {
	Name        string
	Type        string
	Sensitive   bool
	Required    bool
	Default     interface{}
	Constraints []string
	View        view.Node
}

func (d *Descriptor) ToESDL() string {
	var b strings.Builder
	b.WriteString("property ")
	b.WriteString(d.Name)
	b.WriteString(" -> " + d.Type)
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
