package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
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

func (c *Controller) GetFilesHandler(w http.ResponseWriter, r *http.Request) {
	if res, err := json.Marshal(map[string]any{"error": false, "data": c.repo.GetFiles()}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(res))
	}
}

func (c *Controller) UploadHandler(w http.ResponseWriter, r *http.Request) {
	id := uuid.NewString()
	r.ParseMultipartForm(10 << 20)
	rawFile, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Ошибка получения файла: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer rawFile.Close()

	fileName := handler.Filename
	path := r.URL.Query().Get("path")
	dst, err := os.Create(filepath.Join("./data", id))
	if err != nil {
		http.Error(w, "Не удалось сохранить файл: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// 4. Копируем содержимое загруженного файла в созданный файл
	if _, err := io.Copy(dst, rawFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	size := handler.Size
	// save file in database
	if err := c.repo.CreateFile(id, fileName, path, "file", size); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// form response content
	res, err := json.Marshal(map[string]string{"id": id})
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
	if err := c.repo.CreateFile(id, file.Name, file.Path, "folder", 0); err != nil {
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

func (c *Controller) GetFileHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	http.ServeFile(w, r, "data/"+id)
}

func (c *Controller) RenameFileHandler(w http.ResponseWriter, r *http.Request) {
	var file *models.File
	if err := json.NewDecoder(r.Body).Decode(&file); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	newFilename := file.Name

	file, err := c.repo.GetFile(file.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	c.repo.RenameItem(file.ID, newFilename)
	w.WriteHeader(http.StatusNoContent)

	return
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
