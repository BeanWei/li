package lient

import (
	"reflect"

	"entgo.io/ent/schema"
	"github.com/BeanWei/li/li-engine/view/ui"
	"github.com/gogf/gf/v2/text/gstr"
)

type (
	Annotation struct {
		ViewSchema    *ui.Schema   `json:"ViewSchema,omitempty"`
		ColumnProps   *ColumnProps `json:"ColumnProps,omitempty"`
		ValidateRule  string       `json:"ValidateRule,omitempty"`
		DisableCreate bool         `json:"DisableCreate,omitempty"`
		DisableRead   bool         `json:"DisableRead,omitempty"`
		DisableUpdate bool         `json:"DisableUpdate,omitempty"`

		// 解耦的 Edge 关联查询
		EdgeType        string `json:"EdgeType,omitempty"`
		EdgePackage     string `json:"EdgePackage,omitempty"`
		EdgeName        string `json:"EdgeName,omitempty"`
		EdgeStructField string `json:"EdgeStructField,omitempty"`
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

func XEdge(name string, t interface{}) Annotation {
	tn := typ(t)
	return Annotation{
		EdgeType:        tn,
		EdgePackage:     gstr.ToLower(tn),
		EdgeName:        name,
		EdgeStructField: gstr.CaseCamel(name),
	}
}

func typ(t interface{}) string {
	if rt := reflect.TypeOf(t); rt.NumIn() > 0 {
		return rt.In(0).Name()
	}
	return ""
}

var _ schema.Annotation = (*Annotation)(nil)
