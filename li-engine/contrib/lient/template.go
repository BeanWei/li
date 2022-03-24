package lient

import (
	"embed"
	"strings"
	"text/template"

	"entgo.io/ent/entc/gen"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/hexops/valast"
)

var (
	//go:embed template
	templateDir embed.FS
	FuncMap     = template.FuncMap{
		"valast": valast.String,
		"title":  title,
	}
	Templates = gen.MustParse(gen.NewTemplate("lient").Funcs(FuncMap).ParseFS(templateDir, "template/*tmpl"))
)

func title(s string) string {
	return strings.Title(gstr.Join(gstr.SplitAndTrim(gstr.CaseSnake(s), "_"), " "))
}
