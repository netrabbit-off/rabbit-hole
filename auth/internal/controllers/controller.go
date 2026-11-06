package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/netrabbit-off/rabbit-hole/internal/repository"
	"github.com/netrabbit-off/rabbit-hole/pkg/models"
)

type Controller struct {
	repo *repository.Repository
}

func GetController(repo *repository.Repository) *Controller {
	return &Controller{repo: repo}
}

func (c *Controller) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error"))
		return
	}

	id := uuid.NewString()
	c.repo.CreateUser(id, user.Login, user.Password)

	res, _ := json.Marshal(&models.Response{
		Token: "123",
		Error: "false",
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

func (c *Controller) LoginHandler(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error"))
		return
	}

	if c.repo.CheckLogin(user.Login, user.Password) {
		res, _ := json.Marshal(&models.Response{
			Token: "123",
			Error: "false",
		})
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		return
	}

	res, _ := json.Marshal(&models.Response{
		Error:   "invalid_credentials",
		Message: "Неправильный логин или пароль",
	})

	w.WriteHeader(http.StatusUnauthorized)
	w.Write(res)
}
