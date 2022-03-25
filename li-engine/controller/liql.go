package controller

import (
	"context"
	"reflect"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gstructs"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
	"github.com/gogf/gf/v2/util/gvalid"
)

type LiqlReq struct {
	Operation string      `json:"operation" p:"operation" v:"required"`
	Variables interface{} `json:"variables" p:"variables"`
}

func Liql(r *ghttp.Request) {
	var (
		ctx = r.Context()
		req *LiqlReq
		res interface{}
	)
	err := r.ParseForm(&req)
	if err == nil {
		f, exists := handlers[req.Operation]
		if exists {
			var inputValues = []reflect.Value{
				reflect.ValueOf(ctx),
			}
			if f.Type.NumIn() == 2 {
				var (
					inputObject reflect.Value
				)
				if f.Type.In(1).Kind() == reflect.Ptr {
					inputObject = reflect.New(f.Type.In(1).Elem())
					err = doParse(ctx, req.Variables, inputObject.Interface())
				} else {
					inputObject = reflect.New(f.Type.In(1).Elem()).Elem()
					err = doParse(ctx, req.Variables, inputObject.Addr().Interface())
				}
				if err == nil {
					inputValues = append(inputValues, inputObject)
				}
			}
			if len(inputValues) == 2 {
				// Call handler with dynamic created parameter values.
				results := f.Value.Call(inputValues)
				switch len(results) {
				case 1:
					if !results[0].IsNil() {
						if e, ok := results[0].Interface().(error); ok {
							err = e
						}
					}
				case 2:
					res = results[0].Interface()
					if !results[1].IsNil() {
						if e, ok := results[1].Interface().(error); ok {
							err = e
						}
					}
				}
			}
		} else {
			err = gerror.NewCodef(
				gcode.CodeInvalidParameter,
				`parameter operation "%s" is a invalid controller name`,
				req.Operation,
			)
		}
	}

	var (
		code = gerror.Code(err)
		msg  string
	)
	if err != nil {
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()
	} else {
		code = gcode.CodeOK
	}
	internalErr := r.Response.WriteJson(ghttp.DefaultHandlerResponse{
		Code:    code.Code(),
		Message: msg,
		Data:    res,
	})
	if internalErr != nil {
		g.Log().Errorf(ctx, `%+v`, internalErr)
	}
}

func doParse(ctx context.Context, variables interface{}, pointer interface{}) error {
	var (
		reflectVal1  = reflect.ValueOf(pointer)
		reflectKind1 = reflectVal1.Kind()
	)
	if reflectKind1 != reflect.Ptr {
		return gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`invalid parameter type "%v", of which kind should be of *struct/**struct/*[]struct/*[]*struct, but got: "%v"`,
			reflectVal1.Type(),
			reflectKind1,
		)
	}
	var (
		reflectVal2  = reflectVal1.Elem()
		reflectKind2 = reflectVal2.Kind()
	)
	switch reflectKind2 {
	case reflect.Ptr, reflect.Struct:
		data, ok := variables.(map[string]interface{})
		if !ok {
			return gerror.NewCodef(
				gcode.CodeInvalidParameter,
				`parameter variables should be map but got "%T"`,
				variables,
			)
		}
		tagFields, err := gstructs.TagFields(pointer, []string{"d", "default"})
		if err != nil {
			return err
		}
		if len(tagFields) > 0 {
			var (
				foundKey   string
				foundValue interface{}
			)
			for _, field := range tagFields {
				foundKey, foundValue = gutil.MapPossibleItemByKey(data, field.Name())
				if foundKey == "" {
					data[field.Name()] = field.TagValue
				} else {
					if isEmpty(foundValue) {
						data[foundKey] = field.TagValue
					}
				}
			}
		}
		if err = gconv.Struct(data, pointer); err != nil {
			return err
		}
		if err = gvalid.New().
			Bail().
			Data(pointer).
			Assoc(data).
			Run(ctx); err != nil {
			return err
		}
	case reflect.Array, reflect.Slice:
		// TODO: Support default value
		j, err := gjson.LoadContent(variables)
		if err != nil {
			return err
		}
		if err = j.Var().Scan(pointer); err != nil {
			return err
		}
		for i := 0; i < reflectVal2.Len(); i++ {
			if err = gvalid.New().
				Bail().
				Data(reflectVal2.Index(i)).
				Assoc(j.Get(gconv.String(i)).Map()).
				Run(ctx); err != nil {
				return err
			}
		}
	}
	return nil
}
