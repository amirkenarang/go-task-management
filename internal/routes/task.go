package routes

import (
	"example.com/task-managment/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func TaskRoutes(router fiber.Router) {

	router.Get("/tasks", handlers.GetTasks)
	router.Get("/tasks/:id", handlers.GetTask)
	router.Post("/tasks", handlers.CreateTasks)
	router.Put("/tasks/:id", handlers.UpdateTask)
	router.Delete("/tasks/:id", handlers.DeleteTask)
}
