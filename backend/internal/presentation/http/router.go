package http

import (
	"ritmotrack-backend/internal/infrastructure/adapter_di"
	Controllers "ritmotrack-backend/internal/presentation/http/controller"

	apiMiddleware "ritmotrack-backend/internal/presentation/http/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RegisterRouters(r *chi.Mux, container *adapter_di.Container) {
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	container.Invoke(func(
		userController *Controllers.UserController,
		authController *Controllers.AuthController,
		authMiddleware *apiMiddleware.AuthMiddleware,
	) {
		r.Route("/api", func(r chi.Router) {
			r.Get("/users", userController.GetAll)
			r.Get("/users/{id}", userController.GetByID)
			r.Post("/users", userController.Create)

			r.Route("/auth", func(r chi.Router) {
				r.Post("/", authController.CreateJwt)

				r.Route("/check", func(r chi.Router) {

					r.Use(authMiddleware.Middleware)

					r.Get("/", authController.CheckJwt)
				})
			})
		})
	})
}
