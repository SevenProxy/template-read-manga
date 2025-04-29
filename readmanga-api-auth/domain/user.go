package domain

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Passowrd string `json:"password"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}
