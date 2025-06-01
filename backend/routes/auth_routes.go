package routes

import (
	"backend/controllers"

	"github.com/go-chi/chi/v5"
)

func AuthRoutes(r chi.Router) {
    ac := controllers.NewAuthController()
    r.Post("/login", ac.Login)
    r.Post("/register", ac.Register)
}