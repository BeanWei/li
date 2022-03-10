package view

import "github.com/BeanWei/li/li-engine/view/node"

type (
	Schema interface {
		Type() string
		Mixin() []Mixin
		Nodes() []Node
	}

	Mixin interface {
		Nodes() []Node
	}

	Node interface {
		Descriptor() *node.Descriptor
	}
)
