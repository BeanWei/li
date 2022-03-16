package field

import "github.com/BeanWei/li/li-engine/view"

type Descriptor struct {
	Name        string
	Type        string
	Sensitive   bool
	Required    bool
	Default     interface{}
	Constraints []string
	View        view.Node
}
