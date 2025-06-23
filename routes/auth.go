package routes

import (
	"github.com/WelintonJunior/billing-and-subscription-service/controllers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(route fiber.Router) {
	auth := route.Group("/auth")
	auth.Post("/login", controllers.Login())
	auth.Post("/register", controllers.Register())
	auth.Get("/refresh", controllers.RefreshToken())
}
