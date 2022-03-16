package mixin

import "github.com/BeanWei/li/li-engine/model"

type Schema struct{}

func (Schema) Fields() []model.Field { return nil }

func (Schema) Indexes() []model.Index { return nil }
