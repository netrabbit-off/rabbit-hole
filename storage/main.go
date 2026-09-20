package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

type File struct {
	ID      string `json:"id"`
	Path    string `json:"path"`
	Owner   string `json:"owner"`
	Content string `json:"content"`
}

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		res, err := json.Marshal(map[string]string{"message": "xyu"})
		if err != nil {
			return
		}
		w.Write(res)
	})

	r.Post("/upload", func(w http.ResponseWriter, r *http.Request) {
		var file *File
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
	})

	r.Post("/download", func(w http.ResponseWriter, r *http.Request) {
		var file *File
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
	})

	http.ListenAndServe(":8080", r)
}
