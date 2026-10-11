package models

type File struct {
	ID         string `json:"id"`
	Path       string `json:"path"`
	Owner      string `json:"owner"`
	Size       int64  `json:"size"`
	ModifiedAt string `json:"modifiedAt"`
	Content    string `json:"content"`
	Type       string `json:"type"`
}
