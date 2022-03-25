package engine_test

import (
	"testing"

	engine "github.com/BeanWei/li/li-engine"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/node"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
)

type Hello struct {
	view.Schema
}

func (Hello) Nodes() []view.Node {
	return []view.Node{
		node.GridRow("row1").
			Gutter(16).
			Children(
				node.GridCol("col1").
					Span(16),
				node.GridCol("col2").
					Span(8),
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

func Test_GF(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		g.Dump(gconv.Map("") == nil)
		var res *Hello
		gconv.Struct(gconv.Map(""), &res)
		g.Dump(res)
	})
}
