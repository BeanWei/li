package view

import (
	"github.com/BeanWei/li/li-engine/view/ui"
)

type (
	Schema interface {
		Type() string
		Mixin() []Mixin
		Blocks() []Block
	}

	Mixin interface {
		Blocks() []Block
	}

	Block interface {
		Schema() *ui.Schema
	}
)
