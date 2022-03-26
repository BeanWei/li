package dto

type UserSignInReq struct {
	Email    string `v:"required|email"`
	Password string `v:"required"`
}

type UserSignInRes struct {
	ID  string `json:"id"`
	Sid string `json:"sid"`
}
