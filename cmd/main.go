package main

import (
	"example.com/task-managment/internal/db"
	"example.com/task-managment/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	db.InitDB()
	// Create a new Fiber app
	app := fiber.New()

	// Register all routes
	routes.RegisterRoutes(app)

	// Start the server at localhost:8080
	app.Listen(":8080")

}
