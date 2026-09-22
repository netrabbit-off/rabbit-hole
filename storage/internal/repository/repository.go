package repository

import "errors"

type Repository struct {
	DB map[string]map[string]string
}

func NewRepository(db map[string]map[string]string) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) CreateFile(id string, title string, content string) error {
	r.DB[id] = map[string]string{
		"id":      id,
		"title":   title,
		"content": content,
	}
	return nil
}

func (r *Repository) GetFile(id string) (map[string]string, error) {
	if res := r.DB[id]; res != nil {
		return res, nil
	}
	return nil, errors.New("not found")
}
