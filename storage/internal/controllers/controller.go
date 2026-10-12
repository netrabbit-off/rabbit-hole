package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/netrabbit-off/rabbit-hole/internal/repository"
	"github.com/netrabbit-off/rabbit-hole/pkg/helpers"
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
	// form file object
	var file *models.File
	if err := json.NewDecoder(r.Body).Decode(&file); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	file.ID = id
	content := helpers.DecodeB64(file.Content)
	size := int64(len(content))

	// save file in real storage
	if err := helpers.CreateFile(file, content); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	// save file in database
	if err := c.repo.CreateFile(id, file.Path, "file", size); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// form response content
	res, err := json.Marshal(map[string]string{"id": file.ID})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

func (c *Controller) CreateFolderHandler(w http.ResponseWriter, r *http.Request) {
	id := uuid.NewString()
	// form file object
	var file *models.File
	if err := json.NewDecoder(r.Body).Decode(&file); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	file.ID = id

	// save file in database
	if err := c.repo.CreateFile(id, file.Path, "folder", 0); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// form response content
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

	content, err := helpers.ReadFile(file.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	file.Content = helpers.EncodeB64(content)

	res, err := json.Marshal(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(res)
}

func (c *Controller) DeleteFileHandler(w http.ResponseWriter, r *http.Request) {
	var file *models.File
	if err := json.NewDecoder(r.Body).Decode(&file); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	file, err := c.repo.GetFile(file.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	c.repo.DeleteItem(file.ID)
	if file.Type == "folder" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if err := helpers.DeleteFile(file.ID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}
