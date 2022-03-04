package engine

import (
	"testing"

	"github.com/BeanWei/li/li-engine/entity"
	"github.com/BeanWei/li/li-engine/entity/field"
	"github.com/BeanWei/li/li-engine/entity/index"
)

type Post struct {
	entity.Schema
}

func (Post) Mixin() []entity.Mixin {
	return []entity.Mixin{}
}

func (Post) Fields() []entity.Field {
	return []entity.Field{
		field.String("title"),
	}
}

func (Post) Indexes() []entity.Index {
	return []entity.Index{
		index.Fields("title"),
	}
}

func Test_GenEntityESDL(t *testing.T) {
	t.Log("\n" + GenEntityESDL(&Post{}))
}
