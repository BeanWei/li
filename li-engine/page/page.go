package page

import "github.com/BeanWei/li/li-engine/page/node"

type (
	Schema interface {
		Type() string
		Mixin() []Mixin
		Nodes() []Nodes
	}

	Mixin interface {
		Nodes() []Nodes
	}

	Nodes interface {
		Descriptor() *node.Descriptor
	}
)
