package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/netrabbit-off/rabbit-hole/internal/controllers"
)

func main() {
	r := chi.NewRouter()
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		res, err := json.Marshal(map[string]string{"message": "PONG"})
		if err != nil {
			return
		}
		w.Write(res)
	})

	r.Post("/upload", controllers.UploadHandler)
	r.Post("/download", controllers.DownloadHandler)

	http.ListenAndServe(":8080", r)
}
