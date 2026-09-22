package controllers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/netrabbit-off/rabbit-hole/pkg/models"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	var file *models.File
	if err := json.NewDecoder(r.Body).Decode(&file); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	created, err := os.Create("data/" + file.Path)
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

	w.WriteHeader(http.StatusNoContent)
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	var file *models.File
	if err := json.NewDecoder(r.Body).Decode(&file); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	content, err := os.ReadFile("data/" + file.Path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(content))
}
