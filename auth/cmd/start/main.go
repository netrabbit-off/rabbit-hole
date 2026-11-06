package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/netrabbit-off/rabbit-hole/internal/controllers"
	"github.com/netrabbit-off/rabbit-hole/internal/repository"
	"github.com/netrabbit-off/rabbit-hole/pkg/models"
)

func main() {
	r := chi.NewRouter()
	controller := controllers.GetController(repository.GetRepository(map[string]*models.User{}))
	r.Post("/register", controller.RegisterHandler)
	r.Post("/login", controller.LoginHandler)

	http.ListenAndServe(":8089", r)
}
