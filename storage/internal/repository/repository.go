package repository

import (
	"errors"
	"maps"
	"slices"

	"github.com/netrabbit-off/rabbit-hole/pkg/models"
)

type Repository struct {
	DB map[string]*models.File
}

func NewRepository(db map[string]*models.File) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetFiles() []*models.File {
	files := slices.Collect(maps.Values(r.DB))
	return files
}

func (r *Repository) CreateFile(id string, path string, fileType string, size int64) error {
	r.DB[id] = &models.File{
		ID:    id,
		Path:  path,
		Owner: "root",
		Size:  size,
		Type:  fileType,
	}
	return nil
}

func (r *Repository) GetFile(id string) (*models.File, error) {
	if res := r.DB[id]; res != nil {
		return res, nil
	}
	return nil, errors.New("not found")
}

func (r *Repository) DeleteItem(id string) {
	delete(r.DB, id)
}
