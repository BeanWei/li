package view

import (
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/node"
)

type WorkplacePage struct {
	view.Schema
}

func (WorkplacePage) Nodes() []view.Node {
	return []view.Node{
		node.GridRow("row1").
			Gutter(20).
			Children(
				node.GridCol("col1").
					Span(8).
					Content("Col-1"),
				node.GridCol("col2").
					Span(8).
					Content("Col-2"),
				node.GridCol("col3").
					Span(8).
					Content("Col-3"),
			),
		node.GridRow("row2").
			Gutter(20).
			Children(
				node.GridCol("col4").
					Span(16).
					Content("Col-4"),
				node.GridCol("col5").
					Span(8).
					Content("Col-5"),
			),
	}
}
