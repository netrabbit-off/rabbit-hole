package models

type Token struct {
	Hash      string `json:"hash"`
	UserID    string `json:"user_id"`
	IsRefresh bool   `json:"is_refresh"`
}
