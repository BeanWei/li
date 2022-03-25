package controller

import (
	"reflect"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type handler struct {
	Type  reflect.Type
	Value reflect.Value
}

var handlers = make(map[string]*handler)

func Bind(name string, f interface{}) {
	reflectType := reflect.TypeOf(f)
	if reflectType.NumIn() == 0 || reflectType.NumIn() > 2 {
		panic(gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`invalid controller: defined as "%s", but "func(context.Context)" or "func(context.Context, BizRequest)" is required`,
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
	if reflectType.NumIn() == 2 && reflectType.In(1).String() == "interface {}" {
		panic(gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`invalid controller: defined as "%s", but the second input parameter should not be type of "interface {}"`,
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
	handlers[name] = &handler{
		Type:  reflectType,
		Value: reflect.ValueOf(f),
	}
}
