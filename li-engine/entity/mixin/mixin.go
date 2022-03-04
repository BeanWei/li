package mixin

import "github.com/BeanWei/li/li-engine/entity"

type Schema struct{}

func (Schema) Fields() []entity.Field { return nil }

func (Schema) Indexes() []entity.Index { return nil }
