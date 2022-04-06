package view

import (
	"github.com/BeanWei/li/li-app/internal/data/ent"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/node"
)

type BlogTag struct {
	view.Schema
}

func (BlogTag) Nodes() []view.Node {
	return []view.Node{
		node.GridRow("row1").
			Gutter(16).
			Children(
				node.GridCol("col1").
					Span(24).
					Children(
						ent.ListTagView(),
					),
			),
	}
}
