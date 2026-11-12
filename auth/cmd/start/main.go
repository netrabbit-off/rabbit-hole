package main

import (
	"net/http"

	"github.com/go-chi/chi"
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

	userRepo := repository.GetUserRepository(map[string]*models.User{})
	tokenRepo := repository.GetTokenRepository("127.0.0.1:6379", "", 0, 2)
	controller := controllers.GetController(userRepo, tokenRepo)

	r.Post("/update", controller.UpdateTokensHandler)
	r.Post("/register", controller.RegisterHandler)
	r.Get("/check", controller.CheckAuthHandler)
	r.Post("/login", controller.LoginHandler)
	r.Get("/get_me", controller.GetMeHandler)
	// r.Get("/user/{id}", controller.GetUserHandler)

	http.ListenAndServe(":8089", r)
}
