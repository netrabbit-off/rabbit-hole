package models

type User struct {
	ID       string `json:"id"`
	Login    string `json:"user"`
	Password string `json:"password"`
}
