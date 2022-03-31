package engine

import (
	"context"

	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/controller"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/ui"
	"github.com/gogf/gf/v2/container/garray"
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
		AC       ac.AC
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

const (
	OperationGetAppConfig   = "@getAppConfig"
	OperationGetAppView     = "@getAppView"
	OperationGetCurrentUser = "@getCurrentUser"
	OperationUploadFile     = "@uploadFile"
)

func NewApp(cfg *App) {
	var (
		appcfg = app{
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
				for p := range ac.GetAll() {
					if gstr.HasPrefix(p, key) {
						pageacl[key] = append(pageacl[key], p)
					}
				}
				if menu.AC != nil {
					ac.Bind(key, menu.AC)
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
		controller.Bind(OperationGetCurrentUser, cfg.Binding.GetCurrentUserController)
	}

	controller.Bind(OperationGetAppConfig, func(ctx context.Context) (res *app, err error) {
		pageKeys := make([]string, len(pages))
		i := 0
		for pk := range pages {
			pageKeys[i] = pk
			i++
		}
		removes, err := ac.CheckForView(ctx, pageKeys...)
		if err != nil {
			return nil, err
		}
		selfAppCfg := appcfg
		if len(removes) > 0 {
			selfAppCfg.Menus = rebuildAppMenu(selfAppCfg.Menus, garray.NewStrArrayFrom(removes))
		}
		return &selfAppCfg, nil
	})
	controller.Bind(OperationGetAppView, func(ctx context.Context, req *GetAppViewReq) (res *GetAppViewRes, err error) {
		page, exists := pages[req.Key]
		if !exists {
			err = gerror.NewCode(gcode.CodeInvalidParameter, "无效的key")
			return
		}
		res = &GetAppViewRes{
			Schema: page,
		}
		res.Removes, err = ac.CheckForView(ctx, pageacl[req.Key]...)
		return
	})
	controller.Bind(OperationUploadFile, controller.FileUpload)

	s := g.Server()
	s.BindHandler("POST:/api/liql", controller.Liql)
	s.BindHandler("GET:/upload/{bucket_name}/{file_name}", controller.FilePreviw)
}

func rebuildAppMenu(old []*appmenu, removes *garray.StrArray) []*appmenu {
	menus := make([]*appmenu, 0)
	for _, m := range old {
		if !removes.Contains(m.Key) {
			var hasChildren bool
			if len(m.Children) > 0 {
				hasChildren = true
				m.Children = rebuildAppMenu(m.Children, removes)
			}
			if len(m.Children) > 0 || !hasChildren {
				menus = append(menus, m)
			}
		} else if m.Target != "" {
			menus = append(menus, m)
		}
	}
	return menus
}
