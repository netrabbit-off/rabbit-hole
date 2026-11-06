package repository

import (
	"maps"

	"github.com/netrabbit-off/rabbit-hole/pkg/models"
)

type Repository struct {
	DB map[string]*models.User
}

func GetRepository(DB map[string]*models.User) *Repository {
	return &Repository{DB: DB}
}

func (r *Repository) CreateUser(id string, login string, password string) error {
	r.DB[id] = &models.User{
		Login:    login,
		Password: password,
	}
	return nil
}

func (r *Repository) CheckLogin(login string, password string) bool {
	for i := range maps.Keys(r.DB) {
		if r.DB[i].Login == login && r.DB[i].Password == password {
			return true
		}
	}
	return false
}
