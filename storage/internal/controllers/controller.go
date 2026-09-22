package controllers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/netrabbit-off/rabbit-hole/internal/repository"
	"github.com/netrabbit-off/rabbit-hole/pkg/models"
)

type Controller struct {
	repo *repository.Repository
}

func NewController(repo *repository.Repository) *Controller {
	return &Controller{repo: repo}
}

func (c *Controller) UploadHandler(w http.ResponseWriter, r *http.Request) {
	id := uuid.NewString()
	var file *models.File
	if err := json.NewDecoder(r.Body).Decode(&file); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	file.ID = id

	created, err := os.Create("data/" + file.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	_, err = created.Write([]byte(file.Content))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	res, err := json.Marshal(map[string]string{"id": file.ID})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

func (c *Controller) DownloadHandler(w http.ResponseWriter, r *http.Request) {
	var file *models.File
	if err := json.NewDecoder(r.Body).Decode(&file); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	content, err := os.ReadFile("data/" + file.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	file.Content = string(content)

	res, err := json.Marshal(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(res)
}
