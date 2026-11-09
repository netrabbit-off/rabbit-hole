package repository

import (
	"maps"

	"github.com/netrabbit-off/rabbit-hole/pkg/models"
)

type UserRepository struct {
	DB map[string]*models.User
}

func GetUserRepository(DB map[string]*models.User) *UserRepository {
	return &UserRepository{DB: DB}
}

func (r *UserRepository) CreateUser(id string, login string, password string) error {
	r.DB[id] = &models.User{
		Login:    login,
		Password: password,
	}
	return nil
}

func (r *UserRepository) CheckLogin(login string, password string) bool {
	for i := range maps.Keys(r.DB) {
		if r.DB[i].Login == login && r.DB[i].Password == password {
			return true
		}
	}
	return false
}
