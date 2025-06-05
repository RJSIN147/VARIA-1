package routes

import (
	"backend/controllers"

	"github.com/go-chi/chi/v5"
)

func AuthRoutes(r chi.Router) {
	ac := controllers.NewAuthController()
	r.Mount("/auth", ac.Routes()) // Mounts the auth controller routes under /auth
}