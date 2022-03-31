package engine_test

import (
	"testing"

	engine "github.com/BeanWei/li/li-engine"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/node"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/test/gtest"
)

type Hello struct {
	view.Schema
}

func (Hello) Nodes() []view.Node {
	return []view.Node{
		node.GridRow("row_1").
			Gutter(16).
			Children(
				node.GridRow("row_1_1").
					Children(
						node.GridRow("row_1_1_1").
							Children(
								node.GridRow("row_1_1_1_1").
									Children(
										node.GridRow("row_1_1_1_1_1"),
									),
							),
					),
			),
	}
}

func Test_NewApp(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		engine.NewApp(&engine.App{
			Title:     "Li Admin",
			Copyright: "Powered by ❤️璃❤️",
			Menus: []*engine.AppMenu{
				{
					Title: "Welcome",
					Icon:  "IconSmile",
					Children: []*engine.AppMenu{
						{
							Title:  "Hello",
							Page:   new(Hello),
							IsHome: true,
						},
					},
				},
			},
		})
	})
}

func Test_Node(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		_, s := view.ToPage(new(Hello))
		j, _ := gjson.LoadContent(s)
		j.Dump()
	})
}

func Test_GF(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {})
}
