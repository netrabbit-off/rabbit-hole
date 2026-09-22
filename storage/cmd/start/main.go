package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/netrabbit-off/rabbit-hole/internal/controllers"
	"github.com/netrabbit-off/rabbit-hole/internal/repository"
)

func main() {
	r := chi.NewRouter()
	repo := repository.NewRepository(map[string]map[string]string{})
	controller := controllers.NewController(repo)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		res, err := json.Marshal(map[string]string{"message": "PONG"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(res)
	})

	r.Post("/upload", controller.UploadHandler)
	r.Post("/download", controller.DownloadHandler)

	http.ListenAndServe(":8080", r)
}
