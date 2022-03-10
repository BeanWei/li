package engine

import (
	"encoding/json"
	"testing"

	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/node"
)

type PostListPage struct {
	view.Schema
}

func (PostListPage) Mixin() []view.Mixin {
	return []view.Mixin{}
}

func (PostListPage) Nodes() []view.Node {
	return []view.Node{
		node.GridRow("row1").
			Gutter(20).
			Children(
				node.GridCol("col1").
					Span(8).
					Children(
						node.Checkbox("check1"),
					),
				node.GridCol("col2").
					Span(8).
					Children(
						node.Checkbox("check2"),
					),
				node.GridCol("col3").
					Span(8).
					Children(
						node.Checkbox("check3"),
					),
			),
		node.GridRow("row2").
			Gutter(20).
			Children(
				node.GridCol("col4").
					Span(16).
					Children(
						node.Checkbox("check4"),
					),
				node.GridCol("col5").
					Span(8).
					Children(
						node.Checkbox("check5"),
					),
			),
	}
}

func Test_GenPageSchema(t *testing.T) {
	schema := GenPageSchema(&PostListPage{})
	res, _ := json.MarshalIndent(schema, "", "	")
	t.Log("\n" + string(res))
}
