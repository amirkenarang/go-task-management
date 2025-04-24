package routes

import (
	"example.com/task-management/internal/cache"
	"example.com/task-management/internal/db"
	"example.com/task-management/internal/handlers"
	"example.com/task-management/internal/repository"
	"github.com/gofiber/fiber/v2"
)

// AutRoutes registers auth-related endpoints
func AutRoutes(router fiber.Router) {

	userRepo := repository.NewUserRepository(db.DB, cache.RedisClient)
	authHandler := handlers.NewAuthHandler(userRepo)

	router.Post("/signup", authHandler.SignUp)
	router.Post("/login", authHandler.Login)
}
