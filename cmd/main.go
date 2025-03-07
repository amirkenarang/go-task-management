package main

import (
	"log"

	"example.com/task-managment/internal/cache"
	"example.com/task-managment/internal/db"
	"example.com/task-managment/internal/routes"
	"github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"
)

func main() {

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables.")
	}

	db.InitDB()

	// Initialize Redis
	err := cache.InitRedis()
	if err != nil {
		log.Fatalf("Error in connect to redis")
	}
	// Create a new Fiber app
	app := fiber.New()

	// Register all routes
	routes.RegisterRoutes(app)

	// Start the server at localhost:8080
	app.Listen(":8080")

}
