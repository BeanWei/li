package engine

import (
	"reflect"
	"strings"

	"github.com/BeanWei/li/li-engine/entity"
)

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
