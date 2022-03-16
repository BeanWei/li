package engine

import (
	"context"

	"github.com/BeanWei/li/li-engine/ctrl"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	App struct {
		Title string
		Logo  string
		Menus []*AppMenu
	}

	AppMenu struct {
		Name     string
		Children []*AppMenu
		Page     view.Schema
		Target   string
		IsHome   bool
	}

	app struct {
		Title     string     `json:"title"`
		Logo      string     `json:"logo"`
		Menus     []*appmenu `json:"menus"`
		EntryPage string     `json:"entry_page"`
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
			Title: cfg.Title,
			Logo:  cfg.Logo,
			Menus: make([]*appmenu, len(cfg.Menus)),
		}
		pages         = make(map[string]map[string]interface{})
		recursionmenu func(menus []*AppMenu) []*appmenu
	)

	recursionmenu = func(menus []*AppMenu) []*appmenu {
		amenus := make([]*appmenu, len(menus))
		for i, menu := range menus {
			key, page := view.SchemaToMap(menu.Page)
			pages[key] = page
			amenu := &appmenu{
				Name:     menu.Name,
				Key:      key,
				Target:   menu.Target,
				Children: make([]*appmenu, len(menu.Children)),
			}
			if menu.IsHome {
				appcfg.EntryPage = key
			}
			amenu.Children = recursionmenu(menu.Children)
			amenus[i] = amenu
		}
		return amenus
	}

	appcfg.Menus = recursionmenu(cfg.Menus)

	ctrl.RegisterController("@getAppSettings", func(ctx context.Context, variables *gjson.Json) (res interface{}, err error) {
		return appcfg, nil
	})
	ctrl.RegisterController("@getAppPageSchema", func(ctx context.Context, variables *gjson.Json) (res interface{}, err error) {
		return pages[variables.Get("key").String()], nil
	})

	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Bind(ctrl.Liql)
	})
	s.Run()
}
