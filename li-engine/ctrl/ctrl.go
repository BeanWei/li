package ctrl

import (
	"context"

	"github.com/gogf/gf/v2/encoding/gjson"
)

type Controller func(ctx context.Context, variables *gjson.Json) (res interface{}, err error)

var controllers = make(map[string]Controller)

func RegisterController(name string, controller Controller) {
	controllers[name] = controller
}
