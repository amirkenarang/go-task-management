package routes

import (
	"example.com/task-management/internal/db"
	"example.com/task-management/internal/handlers"
	"example.com/task-management/internal/repository"
	"github.com/gofiber/fiber/v2"
)

// UserRoutes registers user-related endpoints
func UserRoutes(router fiber.Router) {

	userRepo := repository.NewUserRepository(db.DB)
	userHandler := handlers.NewUserHandler(userRepo)

	router.Post("/signup", userHandler.SignUp)
	router.Post("/login", userHandler.Login)
}
