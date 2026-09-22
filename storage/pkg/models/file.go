package models

type File struct {
	ID      string `json:"id"`
	Path    string `json:"path"`
	Owner   string `json:"owner"`
	Content string `json:"content"`
}
