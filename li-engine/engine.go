package engine

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/BeanWei/li/li-engine/entity"
	"github.com/edgedb/edgedb-go"
)

func SetupEntity(ctx context.Context, db *edgedb.Client, schemas ...entity.Schema) error {
	var b strings.Builder
	b.WriteString("module default {\n")
	for _, schema := range schemas {
		b.WriteString("	type " + reflect.TypeOf(schema).Elem().Name() + " {\n ")
		for _, field := range schema.Fields() {
			b.WriteString("		" + field.Descriptor().ToESDL())
		}
		b.WriteString("\n	}\n")
	}
	b.WriteString("}")

	err := db.Execute(ctx, fmt.Sprintf(`
START MIGRATION TO {
	%s
};
POPULATE MIGRATION;
COMMIT MIGRATION;
	`, b.String()))
	return err
}
