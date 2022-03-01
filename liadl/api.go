package liadl

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type EdgeqlReq struct {
	Operation string                 `json:"op" v:"required"`
	Args      map[string]interface{} `json:"args" v:"required"`
}

type GetPageReq struct {
	Path string `json:"path" v:"required"`
}

func apisetup() {
	s := g.Server()
	s.Group("/api/liadl", func(group *ghttp.RouterGroup) {
		group.POST("/edgeql", func(ctx context.Context, req *EdgeqlReq) (res []byte, err error) {
			cmd, exists := atm.Commands[req.Operation]
			if !exists {
				return
			}
			cmd, err = g.View().ParseContent(ctx, cmd, req.Args)
			if err != nil {
				return nil, err
			}
			err = db.QuerySingleJSON(ctx, cmd, &res)
			return
		})
		group.GET("/page", func(ctx context.Context, req *GetPageReq) (res map[string]interface{}, err error) {
			page, exists := atm.Pages[req.Path]
			if !exists {
				return
			}
			return page, nil
		})
	})
}
