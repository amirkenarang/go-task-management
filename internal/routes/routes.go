package routes

import (
	"example.com/task-management/internal/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func RegisterRoutes(app *fiber.App) {
	// Enable CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allow all origins (change for production)
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	api := app.Group("/api") // Group all API routes under `/api`

	AutRoutes(api)

	// Task routes needs to authentication, then I add authenticated middlewares to it
	authenticated := api.Group("/") // Group all API routes under `/api`
	authenticated.Use(middlewares.Authenticate)

	TaskRoutes(authenticated)
	UserRoutes(authenticated)

}
