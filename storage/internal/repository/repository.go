package repository

import (
	"errors"

	"github.com/netrabbit-off/rabbit-hole/pkg/models"
)

type Repository struct {
	DB map[string]*models.File
}

func NewRepository(db map[string]*models.File) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) CreateFile(id string, path string, size int64) error {
	r.DB[id] = &models.File{
		ID:    id,
		Path:  path,
		Owner: "root",
		Size:  size,
		Type:  "file",
	}
	return nil
}

func (r *Repository) GetFile(id string) (*models.File, error) {
	if res := r.DB[id]; res != nil {
		return res, nil
	}
	return nil, errors.New("not found")
}
