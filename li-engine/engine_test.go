package engine

import (
	"testing"

	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/node"
	"github.com/gogf/gf/v2/container/garray"
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
		NewApp(&App{
			Title:     "Li Admin",
			Copyright: "Powered by ❤️璃❤️",
			Menus: []*AppMenu{
				{
					Title: "Welcome",
					Icon:  "IconSmile",
					Children: []*AppMenu{
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

func Test_MenuRebuild(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		res := rebuildAppMenu([]*appmenu{
			{Key: "0", Children: []*appmenu{{
				Key: "0-0",
				Children: []*appmenu{
					{
						Key: "0-0-0",
						Children: []*appmenu{
							{
								Key: "0-0-0-0",
							},
						},
					},
				},
			}}},
			{Key: "1", Children: []*appmenu{{
				Key: "1-0",
			}, {
				Key: "1-1",
				Children: []*appmenu{
					{
						Key: "1-1-0",
					},
				},
			}}},
		}, garray.NewStrArrayFrom([]string{"0-0-0", "1-1"}))
		t.AssertEQ(len(res), 1)
		t.AssertEQ(res[0].Key, "1")
		t.AssertEQ(res[0].Children[0].Key, "1-0")
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
