package model

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"

	"ariga.io/atlas/cmd/action"
	"ariga.io/atlas/sql/schema"
	"github.com/gogf/gf/v2/frame/g"
)

func Migrate(ctx context.Context, models ...*Schema) {
	d, err := action.DefaultMux.OpenAtlas(g.DB().GetConfig().Link)
	checkerr(err)
	tables := make([]*schema.Table, len(models))
	for i, m := range models {
		tables[i] = m.table()
	}
	current, err := d.InspectSchema(ctx, "", &schema.InspectOptions{
		Tables: func() (t []string) {
			for i := range models {
				t = append(t, tables[i].Name)
			}
			return t
		}(),
	})
	checkerr(err)
	desired := &schema.Schema{
		Name:   current.Name,
		Attrs:  current.Attrs,
		Tables: tables,
	}
	changes, err := d.SchemaDiff(current, desired)
	checkerr(err)
	if len(changes) == 0 {
		fmt.Println("Schema is synced, no changes to be made")
		return
	}
	p, err := d.PlanChanges(ctx, "plan", changes)
	checkerr(err)
	fmt.Println("-- Planned Changes:")
	for _, c := range p.Changes {
		if c.Comment != "" {
			fmt.Println("--", strings.ToUpper(c.Comment[:1])+c.Comment[1:])
		}
		fmt.Println(c.Cmd)
	}
	if g.DB().GetDryRun() {
		return
	}
	checkerr(d.ApplyChanges(ctx, changes))
}

func (s *Schema) table() *schema.Table {
	ts := &schema.Table{
		Name: func() (tn string) {
			tn = s.Table()
			if tn == "" {
				tn = reflect.TypeOf(s).Elem().Name()
			}
			return
		}(),
	}

	return ts
}

func checkerr(msg interface{}) {
	if msg != nil {
		fmt.Fprintln(os.Stderr, "Error:", msg)
		os.Exit(1)
	}
}
