package lient

import (
	"embed"
	"text/template"

	"entgo.io/ent/entc/gen"
	"github.com/hexops/valast"
)

var (
	//go:embed template
	templateDir embed.FS
	FuncMap     = template.FuncMap{
		"valast": valast.String,
	}
	Templates = gen.MustParse(gen.NewTemplate("lient").Funcs(FuncMap).ParseFS(templateDir, "template/*tmpl"))
)
