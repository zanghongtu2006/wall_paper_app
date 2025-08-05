package model

type User struct {
	BaseModel
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}
