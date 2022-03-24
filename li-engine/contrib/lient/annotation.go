package lient

import (
	"entgo.io/ent/schema"
	"github.com/BeanWei/li/li-engine/view/ui"
)

type (
	Annotation struct {
		ViewSchema    *ui.Schema   `json:"ViewSchema,omitempty"`
		ColumnProps   *ColumnProps `json:"ColumnProps,omitempty"`
		ValidateRule  string       `json:"ValidateRule,omitempty"`
		DisableCreate bool         `json:"DisableCreate,omitempty"`
		DisableRead   bool         `json:"DisableRead,omitempty"`
		DisableUpdate bool         `json:"DisableUpdate,omitempty"`
	}
	ColumnProps struct {
		Width      int    `json:"Width,omitempty"`
		Align      string `json:"Align,omitempty"`
		Ellipsis   bool   `json:"Ellipsis,omitempty"`
		Filterable bool   `json:"Filterable,omitempty"`
		Sortable   bool   `json:"Sortable,omitempty"`
	}
)

func (Annotation) Name() string { return "LiEnt" }

var _ schema.Annotation = (*Annotation)(nil)
