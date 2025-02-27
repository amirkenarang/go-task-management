package routes

import (
	"example.com/task-managment/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api") // Group all API routes under `/api`

	UserRoutes(api)

	// Task routes needs to authentication, then I add authenticated middlewares to it
	authenticated := api.Group("/") // Group all API routes under `/api`
	authenticated.Use(middlewares.Authenticate)

	TaskRoutes(authenticated)

}
