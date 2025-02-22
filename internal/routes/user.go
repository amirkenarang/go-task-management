package routes

import (
	"example.com/task-managment/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

// UserRoutes registers user-related endpoints
func UserRoutes(router fiber.Router) {

	router.Post("/signup", handlers.SignUp)
	router.Post("/login", handlers.Login)
}
