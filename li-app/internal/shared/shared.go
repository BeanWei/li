package shared

// CtxUser 上下文用户信息
type CtxUser struct {
	ID      string `json:"id,omitempty"`
	IsAdmin bool   `json:"is_admin,omitempty"`
}
