package http

import (
	"ritmotrack-backend/internal/core/di"
	Controllers "ritmotrack-backend/internal/presentation/http/controller"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RegisterRouters(r *chi.Mux, container *di.Container) {
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	container.Invoke(func(userController *Controllers.UserController) {
		r.Route("/api", func(r chi.Router) {
			r.Get("/users", userController.GetAll)
			r.Get("/users/{id}", userController.GetByID)
			r.Post("/users", userController.Create)
		})
	})
}
