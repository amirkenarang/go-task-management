package routes

import (
	"example.com/task-managment/internal/db"
	"example.com/task-managment/internal/handlers"
	"example.com/task-managment/internal/repository"
	"github.com/gofiber/fiber/v2"
)

// UserRoutes registers user-related endpoints
func UserRoutes(router fiber.Router) {

	userRepo := repository.NewUserRepository(db.DB)
	userHandler := handlers.NewUserHandler(userRepo)

	router.Post("/signup", userHandler.SignUp)
	router.Post("/login", userHandler.Login)
}
