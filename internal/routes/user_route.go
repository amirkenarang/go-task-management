package routes

import (
	"example.com/task-management/internal/cache"
	"example.com/task-management/internal/db"
	"example.com/task-management/internal/handlers"
	"example.com/task-management/internal/repository"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router) {

	userRepo := repository.NewUserRepository(db.DB, cache.RedisClient)
	userHandler := handlers.NewUserHandler(userRepo)

	router.Post("/users", userHandler.CreateUser)
	router.Get("/users", userHandler.GetUsers)
	router.Get("/users/:id", userHandler.GetUser)
	router.Put("/users/:id", userHandler.UpdateUser)
	router.Delete("/users/:id", userHandler.DeleteUser)
}
