package dto

import "github.com/BeanWei/li/li-app/internal/data/ent"

type UserProfileReq struct{}

type UserProfileRes struct {
	*ent.User
}

type UserSignInReq struct {
	Passport string `v:"required"`
	Password string `v:"required"`
}

type UserSignInRes struct {
	ID  string `json:"id"`
	Sid string `json:"sid"`
}
