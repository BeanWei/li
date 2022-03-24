package lient

import (
	"entgo.io/ent/schema"
	"github.com/BeanWei/li/li-engine/view/ui"
)

type (
	Annotation struct {
		ViewSchema    *ui.Schema
		ColumnProps   *ColumnProps
		ValidateRule  string
		DisableCreate bool
		DisableRead   bool
		DisableUpdate bool
	}
	ColumnProps struct {
		Title      string `json:"title,omitempty"`
		Width      int    `json:"width,omitempty"`
		Align      string `json:"align,omitempty"`
		Ellipsis   bool   `json:"ellipsis,omitempty"`
		Filterable bool
		Sortable   bool
	}
)

func (Annotation) Name() string { return "LiEnt" }

var _ schema.Annotation = (*Annotation)(nil)
