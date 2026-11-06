package models

type Response struct {
	Token   string `json:"token"`
	Error   string `json:"error"`
	Message string `json:"message"`
}
