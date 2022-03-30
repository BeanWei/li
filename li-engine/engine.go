package engine

import (
	"context"
	"sync"

	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/controller"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/ui"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
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
	GetAppViewRes struct {
		Schema  string   `json:"schema"`
		Removes []string `json:"removes"`
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
		pages         = make(map[string]string)
		pageacl       = make(map[string][]string)
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
				var (
					wg      sync.WaitGroup
					acpaths = make(chan string)
					acl     = ac.GetAll()
				)
				wg.Add(len(acl))
				for k := range acl {
					k := k
					go func() {
						defer wg.Done()
						if gstr.Contains(page, k) {
							acpaths <- key
						}
					}()
				}
				go func() {
					wg.Wait()
					close(acpaths)
				}()
				pageacl[key] = make([]string, 0)
				for p := range acpaths {
					pageacl[key] = append(pageacl[key], p)
				}

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
	controller.Bind("@getAppView", func(ctx context.Context, req *GetAppViewReq) (res *GetAppViewRes, err error) {
		page, exists := pages[req.Key]
		if !exists {
			err = gerror.NewCode(gcode.CodeInvalidParameter, "无效的key")
		}
		res = &GetAppViewRes{
			Schema: page,
		}
		res.Removes, err = ac.CheckForView(ctx, pageacl[req.Key]...)
		return
	})
	controller.Bind("@uploadFile", controller.FileUpload)

	s := g.Server()
	s.BindHandler("POST:/api/liql", controller.Liql)
	s.BindHandler("GET:/upload/{bucket_name}/{file_name}", controller.FilePreviw)
}
