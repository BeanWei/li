package engine

import (
	"context"

	"github.com/BeanWei/li/li-engine/control"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/ui"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	App struct {
		Title     string
		Logo      string
		Copyright string
		Menus     []*AppMenu
		NavItems  []view.Node
		Binding   *AppBinding
	}

	AppMenu struct {
		Title    string
		Icon     string
		Page     view.Schema
		Target   string
		IsHome   bool
		Children []*AppMenu
	}

	AppBinding struct {
		SignPage                 view.Schema
		GetCurrentUserController control.Controller
	}

	app struct {
		Title     string       `json:"title"`
		Logo      string       `json:"logo"`
		Copyright string       `json:"copyright"`
		NavItems  []*ui.Schema `json:"navitems"`
		Menus     []*appmenu   `json:"menus"`
		Home      string       `json:"home"`
		Binding   *appbinding  `json:"binding"`
	}

	appmenu struct {
		Key      string     `json:"key"`
		Title    string     `json:"title"`
		Icon     string     `json:"icon"`
		Target   string     `json:"target"`
		Children []*appmenu `json:"children"`
	}

	appbinding struct {
		SignPage map[string]interface{} `json:"signpage"`
	}
)

func NewApp(cfg *App) {
	var (
		_, signpage = view.ToPage(cfg.Binding.SignPage)
		appcfg      = &app{
			Title:     cfg.Title,
			Logo:      cfg.Logo,
			Copyright: cfg.Copyright,
			Menus:     make([]*appmenu, len(cfg.Menus)),
			NavItems:  make([]*ui.Schema, len(cfg.NavItems)),
			Binding: &appbinding{
				SignPage: signpage,
			},
		}
		pages         = make(map[string]map[string]interface{})
		recursionmenu func(menus []*AppMenu) []*appmenu
	)

	recursionmenu = func(menus []*AppMenu) []*appmenu {
		amenus := make([]*appmenu, len(menus))
		for i, menu := range menus {
			key, page := view.ToPage(menu.Page)
			if key != "" {
				pages[key] = page
			}
			amenu := &appmenu{
				Key:      key,
				Title:    menu.Title,
				Icon:     menu.Icon,
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

	for i, ni := range cfg.NavItems {
		appcfg.NavItems[i] = ni.Schema()
	}

	control.RegisterController("@getCurrentUser", cfg.Binding.GetCurrentUserController)
	control.RegisterController("@getAppConfig", func(ctx context.Context, variables *gjson.Json) (res interface{}, err error) {
		return appcfg, nil
	})
	control.RegisterController("@getAppView", func(ctx context.Context, variables *gjson.Json) (res interface{}, err error) {
		return pages[variables.Get("key").String()], nil
	})

	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Bind(control.Liql)
	})
}
