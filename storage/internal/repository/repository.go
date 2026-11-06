package repository

import (
	"errors"
	"log"
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

func (r *Repository) CreateFile(id string, fileName string, path string, fileType string, size int64) error {
	r.DB[id] = &models.File{
		ID:    id,
		Name:  fileName,
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
	item := r.DB[id]
	delete(r.DB, id)
	if item.Type == "file" {
		return
	}

	files := r.GetFiles()
	for _, i := range files {
		if i.Path == item.Path+item.Name+"/" {
			r.DeleteItem(i.ID)
		}
	}
}

func (r *Repository) RenameItem(id string, newFilename string) {
	log.Println("RenameItem called:", id, newFilename)
	item := r.DB[id]
	itemName := item.Name
	r.DB[id].Name = newFilename
	if item.Type == "file" {
		return
	}

	files := r.GetFiles()
	log.Println("  files count:", len(files))
	for _, i := range files {
		log.Println(i.Path, item.Path+itemName+"/", item.Path+newFilename+"/")
		if i.Path == item.Path+itemName+"/" {
			r.MoveItem(i.ID, item.Path+newFilename+"/")
		}
	}
}

func (r *Repository) MoveItem(id string, newPath string) {
	log.Println("MoveItem called:", id, newPath)
	item := r.DB[id]
	itemPath := item.Path
	r.DB[id].Path = newPath
	if item.Type == "file" {
		return
	}

	files := r.GetFiles()
	log.Println("  files count:", len(files))
	for _, i := range files {
		log.Println(i.Path, itemPath+item.Name+"/", newPath+item.Name+"/")
		if i.Path == itemPath+item.Name+"/" {
			r.MoveItem(i.ID, newPath+item.Name+"/")
		}
	}
}
