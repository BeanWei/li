package lient

import (
	"embed"
	"strings"
	"text/template"

	"entgo.io/ent/entc/gen"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/hexops/valast"
)

var (
	//go:embed template
	templateDir embed.FS
	FuncMap     = template.FuncMap{
		"valast":    valast.String,
		"title":     title,
		"contains":  contains,
		"listfield": listfield,
	}
	Templates = gen.MustParse(gen.NewTemplate("lient").Funcs(FuncMap).ParseFS(templateDir, "template/*tmpl"))
)

func title(s string) string {
	return strings.Title(gstr.Join(gstr.SplitAndTrim(gstr.CaseSnake(s), "_"), " "))
}

func contains(s1 interface{}, s2 string) bool {
	return gstr.Contains(gconv.String(s1), s2)
}

func listfield(fields ...*gen.Field) []*gen.Field {
	return fields
}
