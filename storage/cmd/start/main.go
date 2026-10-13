package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/netrabbit-off/rabbit-hole/internal/controllers"
	"github.com/netrabbit-off/rabbit-hole/internal/repository"
	"github.com/netrabbit-off/rabbit-hole/pkg/models"
)

func main() {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
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

	r.Get("/files", controller.GetFilesHandler)
	r.Post("/upload", controller.UploadHandler)
	r.Post("/delete", controller.DeleteFileHandler)
	r.Post("/download", controller.DownloadHandler)
	r.Post("/create_folder", controller.CreateFolderHandler)

	http.ListenAndServe(":8080", r)
}
