package engine

import (
	"encoding/json"
	"testing"

	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/block"
	"github.com/BeanWei/li/li-engine/view/node"
)

type PostListPage struct {
	view.Schema
}

func (PostListPage) Mixin() []view.Mixin {
	return []view.Mixin{}
}

func (PostListPage) Blocks() []view.Block {
	return []view.Block{
		block.GridRow("row1").
			Gutter(20).
			Children(
				block.GridCol("col1").
					Span(8).
					Children(
						node.Checkbox("check1"),
					),
				block.GridCol("col2").
					Span(8).
					Children(
						node.Checkbox("check2"),
					),
				block.GridCol("col3").
					Span(8).
					Children(
						node.Checkbox("check3"),
					),
			),
		block.GridRow("row2").
			Gutter(20).
			Children(
				block.GridCol("col4").
					Span(16).
					Children(
						node.Checkbox("check4"),
					),
				block.GridCol("col5").
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
