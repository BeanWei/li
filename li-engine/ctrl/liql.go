package ctrl

import (
	"context"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var Liql = cLiql{}

type cLiql struct{}

type LiqlReq struct {
	g.Meta    `path:"/api/liql" method:"post"`
	Operation string      `json:"operation" v:"required"`
	Variables *gjson.Json `json:"variables"`
}

func (c *cLiql) Liql(ctx context.Context, req *LiqlReq) (res interface{}, err error) {
	controller, exists := controllers[req.Operation]
	if !exists {
		return nil, gerror.NewCodef(gcode.CodeInvalidParameter, `parameter operation "%s" is a invalid controller name`, req.Operation)
	}
	return controller(ctx, req.Variables)
}
