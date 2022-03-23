package lient

import (
	"entgo.io/ent/schema"
	"github.com/BeanWei/li/li-engine/view/ui"
)

type Annotation struct {
	View *ui.Schema `json:"View,omitempty"`
}

func (Annotation) Name() string { return "LiEnt" }

var _ schema.Annotation = (*Annotation)(nil)
