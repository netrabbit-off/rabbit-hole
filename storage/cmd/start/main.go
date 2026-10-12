package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/netrabbit-off/rabbit-hole/internal/controllers"
	"github.com/netrabbit-off/rabbit-hole/internal/repository"
	"github.com/netrabbit-off/rabbit-hole/pkg/models"
)

func main() {
	r := chi.NewRouter()
	repo := repository.NewRepository(map[string]*models.File{})
	controller := controllers.NewController(repo)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		res, err := json.Marshal(repo.DB)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(res)
	})

	r.Post("/upload", controller.UploadHandler)
	r.Post("/delete", controller.DeleteFileHandler)
	r.Post("/download", controller.DownloadHandler)
	r.Post("/create_folder", controller.CreateFolderHandler)

	http.ListenAndServe(":8080", r)
}
