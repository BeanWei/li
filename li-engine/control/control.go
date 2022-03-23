package control

import (
	"reflect"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type controller struct {
	Type  reflect.Type
	Value reflect.Value
}

var controllers = make(map[string]*controller)

func RegisterController(name string, f interface{}) {
	reflectType := reflect.TypeOf(f)
	if reflectType.NumIn() != 2 {
		panic(gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`invalid controller: defined as "%s", but "func(context.Context, BizRequest)" is required`,
			reflectType.String(),
		))
	}
	if reflectType.In(0).String() != "context.Context" {
		panic(gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`invalid controller: defined as "%s", but the first input parameter should be type of "context.Context"`,
			reflectType.String(),
		))
	}
	if (reflectType.NumOut() == 1 && reflectType.Out(0).String() != "error") || reflectType.NumOut() == 2 && reflectType.Out(1).String() != "error" {
		panic(gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`invalid controller: defined as "%s", but the last output parameter should be type of "error"`,
			reflectType.String(),
		))
	}
	controllers[name] = &controller{
		Type:  reflectType,
		Value: reflect.ValueOf(f),
	}
}
