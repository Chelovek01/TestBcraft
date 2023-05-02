package entity

type User struct {
	Id       int    `json:"id"`
	Username string `json:"user_name"`
	Password string `json:"password"`
}
