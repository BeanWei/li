package view

import (
	"github.com/BeanWei/li/li-app/internal/app/admin/model"
	emodel "github.com/BeanWei/li/li-engine/model"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/node"
)

type SystemUser struct {
	view.Schema
}

func (SystemUser) Nodes() []view.Node {
	return []view.Node{
		node.GridRow("row1").
			Gutter(16).
			Children(
				node.GridCol("col1").
					Span(16).
					Children(
						emodel.ToListNode(new(model.User)),
					),
				node.GridCol("col2").
					Span(8).
					Children(
						emodel.ToFormNode(new(model.User)),
					),
			),
	}
}
