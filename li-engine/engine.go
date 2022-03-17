package engine

import (
	"context"

	"github.com/BeanWei/li/li-engine/control"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	App struct {
		Title    string
		Logo     string
		Menus    []*AppMenu
		NavItems []view.Schema
	}

	AppMenu struct {
		Name     string
		Children []*AppMenu
		Page     view.Schema
		Target   string
		IsHome   bool
	}

	app struct {
		Title    string        `json:"title"`
		Logo     string        `json:"logo"`
		NavItems []view.Schema `json:"navitems"`
		Menus    []*appmenu    `json:"menus"`
		Home     string        `json:"home"`
	}

	appmenu struct {
		Name     string     `json:"name"`
		Key      string     `json:"key"`
		Target   string     `json:"target"`
		Children []*appmenu `json:"children"`
	}
)

func NewApp(cfg *App) {
	var (
		appcfg = &app{
			Title:    cfg.Title,
			Logo:     cfg.Logo,
			NavItems: cfg.NavItems,
			Menus:    make([]*appmenu, len(cfg.Menus)),
		}
		pages         = make(map[string]map[string]interface{})
		recursionmenu func(menus []*AppMenu) []*appmenu
	)

	recursionmenu = func(menus []*AppMenu) []*appmenu {
		amenus := make([]*appmenu, len(menus))
		for i, menu := range menus {
			key, page := view.ToPage(menu.Page)
			pages[key] = page
			amenu := &appmenu{
				Name:     menu.Name,
				Key:      key,
				Target:   menu.Target,
				Children: make([]*appmenu, len(menu.Children)),
			}
			if menu.IsHome {
				appcfg.Home = key
			}
			amenu.Children = recursionmenu(menu.Children)
			amenus[i] = amenu
		}
		return amenus
	}

	appcfg.Menus = recursionmenu(cfg.Menus)

	control.RegisterController("@getAppConfig", func(ctx context.Context, variables *gjson.Json) (res interface{}, err error) {
		return appcfg, nil
	})
	control.RegisterController("@getAppView", func(ctx context.Context, variables *gjson.Json) (res interface{}, err error) {
		return pages[variables.Get("uid").String()], nil
	})

	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Bind(control.Liql)
	})
}
