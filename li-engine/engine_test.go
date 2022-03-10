package engine

import (
	"encoding/json"
	"testing"

	"github.com/BeanWei/li/li-engine/entity"
	"github.com/BeanWei/li/li-engine/entity/field"
	"github.com/BeanWei/li/li-engine/entity/index"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/node"
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

type PostListPage struct {
	view.Schema
}

func (PostListPage) Mixin() []view.Mixin {
	return []view.Mixin{}
}

func (PostListPage) Blocks() []view.Block {
	return []view.Block{
		node.Checkbox("a"),
	}
}

func Test_GenPageSchema(t *testing.T) {
	schema := GenPageSchema(&PostListPage{})
	res, _ := json.MarshalIndent(schema, "", "	")
	t.Log("\n" + string(res))
}
