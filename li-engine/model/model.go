package model

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/BeanWei/li/li-engine/model/edge"
	"github.com/BeanWei/li/li-engine/model/field"
	"github.com/BeanWei/li/li-engine/model/index"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/node"
	"github.com/gogf/gf/v2/encoding/gjson"
)

type (
	Schema interface {
		Type()
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

// ToDbSchema 生成模型的 esdl
// Note: 需要注意 schemas 的顺序，关联的对象必须放在前面
func ToDbSchema(schemas ...Schema) string {
	var (
		b          strings.Builder
		linkfields = make(map[string]bool)
	)
	b.WriteString("module default {\n")
	for _, schema := range schemas {
		b.WriteString("	type " + reflect.TypeOf(schema).Elem().Name() + " {\n ")
		// property
		fields := make([]Field, 0)
		for _, mixin := range schema.Mixin() {
			fields = append(fields, mixin.Fields()...)
		}
		fields = append(fields, schema.Fields()...)
		for _, field := range fields {
			b.WriteString("		" + field.Descriptor().ToESDL())
		}
		// property link
		edges := make([]Edge, 0)
		for _, mixin := range schema.Mixin() {
			edges = append(edges, mixin.Edges()...)
		}
		edges = append(edges, schema.Edges()...)
		for _, edge := range edges {
			desc := edge.Descriptor()
			linkfields[desc.Name] = true
			b.WriteString("		" + desc.ToESDL())
		}
		// index
		idxs := make([]Index, 0)
		for _, mixin := range schema.Mixin() {
			idxs = append(idxs, mixin.Indexes()...)
		}
		idxs = append(idxs, schema.Indexes()...)
		for _, idx := range idxs {
			fl := len(idx.Descriptor().Fields)
			idxstrs := make([]string, fl)
			for i, f := range idx.Descriptor().Fields {
				if _, exists := linkfields[f]; exists {
					idxstrs[i] = "__subject__@" + f
				} else {
					idxstrs[i] = "." + f
				}
			}
			if fl == 1 {
				b.WriteString("\n		index on (" + strings.Join(idxstrs, ", ") + ");")
			} else if fl > 1 {
				b.WriteString("\n		index on ((" + strings.Join(idxstrs, ", ") + "));")
			}
		}
		b.WriteString("\n	}\n")
	}
	b.WriteString("}")
	/*
		START MIGRATION TO {
			%s
		};
		POPULATE MIGRATION;
		COMMIT MIGRATION;
	*/
	return b.String()
}

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
		Child(nodes...)
}

func ToListNode(schema Schema) view.Node {
	var (
		entity  = reflect.TypeOf(schema).Elem().Name()
		fields  = make([]Field, 0)
		columns = make([]view.Node, 0)
	)
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
				node.Space("actions").Child(
					node.ListActionRecordEditDrawer("edit").
						Child(columns...).
						ForInit("@get"+entity, func(ctx context.Context, variables *gjson.Json) (res interface{}, err error) {
							return
						}).
						ForSubmit("@update"+entity, func(ctx context.Context, variables *gjson.Json) (res interface{}, err error) {
							return
						}),
					node.ListActionRecordDelete("delete").
						ForSubmit("@delete"+entity, func(ctx context.Context, variables *gjson.Json) (res interface{}, err error) {
							return
						}),
				),
			),
	)

	return node.List(entity).
		Child(
			node.ListAction("actions").Child(
				node.ListActionRecordEditDrawer("add"+entity).
					Child(columns...).
					ForSubmit("@add"+entity, func(ctx context.Context, variables *gjson.Json) (res interface{}, err error) {
						return
					}),
				node.ListActionRowSelection("deleteMany"+entity).
					ForSubmit("@deleteMany"+entity, func(ctx context.Context, variables *gjson.Json) (res interface{}, err error) {
						return
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
