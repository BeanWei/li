package control

import (
	"context"

	"github.com/gogf/gf/v2/encoding/gjson"
)

type (
	Controller         func(ctx context.Context, variables *gjson.Json) (res interface{}, err error)
	ControllerProvider struct {
		Controller Controller
		Validator  interface{}
	}
)

var controllers = make(map[string]*ControllerProvider)

func RegisterController(name string, controller Controller, validators ...interface{}) {
	ctl := &ControllerProvider{
		Controller: controller,
	}
	if len(validators) > 0 {
		ctl.Validator = validators[0]
	} else {
		ctl.Validator = nil
	}
	controllers[name] = ctl
}
