package routes

import (
	"github.com/WelintonJunior/billing-and-subscription-service/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(route fiber.Router) {
	userGroup := route.Group("/users")
	userGroup.Get("/", controllers.ListUsers())
	userGroup.Get("/:id", controllers.GetUserById())
	userGroup.Put("/:id", controllers.UpdateUserById())
	userGroup.Delete("/:id", controllers.DeleteUserById())
}
