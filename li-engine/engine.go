package engine

import (
	"reflect"
	"strings"

	"github.com/BeanWei/li/li-engine/entity"
	"github.com/BeanWei/li/li-engine/view"
)

// GenEntityESDL 生成模型的 esdl
// Note: 需要注意 schemas 的顺序，关联的对象必须放在前面
func GenEntityESDL(schemas ...entity.Schema) string {
	var (
		b          strings.Builder
		linkfields = make(map[string]bool)
	)
	b.WriteString("module default {\n")
	for _, schema := range schemas {
		b.WriteString("	type " + reflect.TypeOf(schema).Elem().Name() + " {\n ")
		// property
		fields := make([]entity.Field, 0)
		for _, mixin := range schema.Mixin() {
			fields = append(fields, mixin.Fields()...)
		}
		fields = append(fields, schema.Fields()...)
		for _, field := range fields {
			desc := field.Descriptor()
			b.WriteString("		" + desc.ToESDL())
			if field.Descriptor().Link.IsLink {
				linkfields[desc.Name] = true
			}
		}
		// index
		idxs := make([]entity.Index, 0)
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

// GenPageSchema 生成页面的 json-schema
func GenPageSchema(schemas ...view.Schema) map[string]map[string]interface{} {
	pages := make(map[string]map[string]interface{})
	for _, schema := range schemas {
		var (
			properties = make(map[string]interface{})
			nodes      = make([]view.Node, 0)
		)
		for _, mixin := range schema.Mixin() {
			nodes = append(nodes, mixin.Nodes()...)
		}
		nodes = append(nodes, schema.Nodes()...)
		for _, node := range nodes {
			properties[node.Schema().Name] = node.Schema()
		}
		pages[reflect.TypeOf(schema).Elem().Name()] = map[string]interface{}{
			"type":       "object",
			"properties": properties,
		}
	}
	return pages
}

// GenApp 生成应用信息
func GenAppConfig() {

}
