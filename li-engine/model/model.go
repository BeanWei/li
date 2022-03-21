package model

import (
	"context"
	"fmt"
	"reflect"

	"github.com/BeanWei/li/li-engine/model/edge"
	"github.com/BeanWei/li/li-engine/model/field"
	"github.com/BeanWei/li/li-engine/model/index"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/node"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

type (
	Schema interface {
		Table() string
		Mixin() []Mixin
		Fields() []Field
		Edges() []Edge
		Indexes() []Index
	}

	Mixin interface {
		Fields() []Field
		Edges() []Edge
		Indexes() []Index
	}

	Field interface {
		Descriptor() *field.Descriptor
	}

	Edge interface {
		Descriptor() *edge.Descriptor
	}

	Index interface {
		Descriptor() *index.Descriptor
	}
)

func ToFormNode(schema Schema) view.Node {
	nodes := make([]view.Node, 0)
	for _, mixin := range schema.Mixin() {
		for _, mfield := range mixin.Fields() {
			nodes = append(nodes, mfield.Descriptor().View)
		}
	}
	for _, field := range schema.Fields() {
		nodes = append(nodes, field.Descriptor().View)
	}
	return node.Form(reflect.TypeOf(schema).Elem().Name()).
		Children(nodes...)
}

func ToListNode(schema Schema) view.Node {
	var (
		table   = schema.Table()
		entity  = reflect.TypeOf(schema).Elem().Name()
		fields  = make([]Field, 0)
		columns = make([]view.Node, 0)
	)
	if table == "" {
		table = gstr.CaseSnake(entity)
	}
	for _, mixin := range schema.Mixin() {
		fields = append(fields, mixin.Fields()...)
	}
	fields = append(fields, schema.Fields()...)

	for i, field := range fields {
		columns = append(
			columns,
			node.ListTableColumn(fmt.Sprintf("column%d", i)).
				Title("").
				DataIndex(field.Descriptor().Name).
				Render(field.Descriptor().View),
		)
	}
	columns = append(
		columns,
		node.ListTableColumn(fmt.Sprintf("column%d", len(fields))).
			Title("操作").
			DataIndex("__action").
			Render(
				node.Space("actions").Children(
					node.ListActionRecordEditDrawer("edit").
						Children(columns...).
						ForInit("@get"+entity, func(ctx context.Context, variables *gjson.Json) (res interface{}, err error) {
							return g.Model(table).Ctx(ctx).One("id", variables.Get("id").String())
						}).
						ForSubmit("@update"+entity, func(ctx context.Context, variables *gjson.Json) (res interface{}, err error) {
							return g.Model(table).Ctx(ctx).Data(variables.Map()).Where("id", variables.Get("id").String()).Update()
						}),
					node.ListActionRecordDelete("delete").
						ForSubmit("@delete"+entity, func(ctx context.Context, variables *gjson.Json) (res interface{}, err error) {
							return g.Model(table).Ctx(ctx).Delete("id", variables.Get("id").String())
						}),
				),
			),
	)

	return node.List(entity).
		ForInit("@list"+entity, func(ctx context.Context, variables *gjson.Json) (res interface{}, err error) {
			// TODO: 完善列表请求
			total, err := g.Model(table).Ctx(ctx).Count()
			if err != nil {
				return nil, err
			} else if total == 0 {
				return
			}
			data, err := g.Model(table).Ctx(ctx).Page(0, 20).All()
			if err != nil {
				return nil, err
			}
			return g.Map{
				"total": total,
				"data":  data,
			}, nil
		}).
		Children(
			node.ListAction("actions").Children(
				node.ListActionRecordEditDrawer("add"+entity).
					Title("新建").
					Children(columns...).
					ForSubmit("@add"+entity, func(ctx context.Context, variables *gjson.Json) (res interface{}, err error) {
						return g.Model(table).Ctx(ctx).Insert(variables.Map())
					}),
				node.ListActionRowSelection("deleteMany"+entity).
					Title("批量删除").
					ForSubmit("@deleteMany"+entity, func(ctx context.Context, variables *gjson.Json) (res interface{}, err error) {
						return g.Model(table).Ctx(ctx).Delete("id IN (?)", variables.Get("ids"))
					}).
					AfterReload().
					ConfirmTitle("确认删除").
					ButtonStatus("danger"),
			),
			node.ListTable("table").
				RowSelectionType("checkbox").
				Columns(columns...),
		)
}
