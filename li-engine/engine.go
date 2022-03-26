package engine

import (
	"context"

	"github.com/BeanWei/li/li-engine/controller"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/ui"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
)

type (
	App struct {
		Title     string
		Logo      string
		Copyright string
		Entry     string
		Menus     []*AppMenu
		NavItems  []view.Node
		Binding   *AppBinding
	}

	AppMenu struct {
		Key      string
		Title    string
		Icon     string
		Page     view.Interface
		Target   string
		IsHome   bool
		Children []*AppMenu
	}

	AppBinding struct {
		SignForm                 view.Node
		GetCurrentUserController interface{}
	}

	app struct {
		Title     string       `json:"title"`
		Logo      string       `json:"logo"`
		Copyright string       `json:"copyright"`
		Entry     string       `json:"entry"`
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
		SignForm *ui.Schema `json:"signform"`
	}

	GetAppViewReq struct {
		Key string `p:"key" v:"required"`
	}
)

func NewApp(cfg *App) {
	var (
		appcfg = &app{
			Title:     cfg.Title,
			Logo:      cfg.Logo,
			Copyright: cfg.Copyright,
			Menus:     make([]*appmenu, len(cfg.Menus)),
			NavItems:  make([]*ui.Schema, len(cfg.NavItems)),
			Entry:     cfg.Entry,
		}
		pages         = make(map[string]map[string]interface{})
		recursionmenu func(menus []*AppMenu) []*appmenu
	)

	recursionmenu = func(menus []*AppMenu) []*appmenu {
		amenus := make([]*appmenu, len(menus))
		for i, menu := range menus {
			key, page := view.ToPage(menu.Page)
			if menu.Key != "" {
				key = menu.Key
			}
			if key != "" {
				pages[key] = page
			} else {
				key = grand.S(8)
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

	if cfg.Binding != nil {
		appcfg.Binding = &appbinding{
			SignForm: cfg.Binding.SignForm.Schema(),
		}
		controller.Bind("@getCurrentUser", cfg.Binding.GetCurrentUserController)
	}

	controller.Bind("@getAppConfig", func(ctx context.Context) (res *app, err error) {
		return appcfg, nil
	})
	controller.Bind("@getAppView", func(ctx context.Context, req *GetAppViewReq) (res map[string]interface{}, err error) {
		return pages[req.Key], nil
	})

	g.Server().BindHandler("POST:/api/liql", controller.Liql)
}
