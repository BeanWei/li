package model

import (
	"github.com/BeanWei/li/li-engine/model"
	"github.com/BeanWei/li/li-engine/model/field"
)

type User struct {
	model.Schema
}

func (User) Table() string {
	return "users"
}

func (User) Fields() []model.Field {
	return []model.Field{
		field.String("email").Unique().Required(),
		field.String("nickname").Required(),
		field.String("password").Sensitive(),
	}
}
