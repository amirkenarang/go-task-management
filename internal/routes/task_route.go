package routes

import (
	"example.com/task-management/internal/cache"
	"example.com/task-management/internal/db"
	"example.com/task-management/internal/handlers"
	"example.com/task-management/internal/repository"
	"github.com/gofiber/fiber/v2"
)

func TaskRoutes(router fiber.Router) {

	taskRepo := repository.NewTaskRepository(db.DB, cache.RedisClient)
	taskHandler := handlers.NewTaskHandler(taskRepo)

	router.Post("/tasks", taskHandler.CreateTasks)
	router.Get("/tasks", taskHandler.GetTasks)
	router.Get("/tasks/:id", taskHandler.GetTask)
	router.Put("/tasks/:id", taskHandler.UpdateTask)
	router.Delete("/tasks/:id", taskHandler.DeleteTask)
}
