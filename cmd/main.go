package main

import (
	"log"

	"example.com/task-management/internal/cache"
	"example.com/task-management/internal/db"
	"example.com/task-management/internal/middlewares"
	"example.com/task-management/internal/monitoring"
	"example.com/task-management/internal/routes"
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

	// Initialize Prometheus
	monitoring.Init()

	// Create a new Fiber app
	app := fiber.New()

	// Add Prometheus middleware
	app.Use(middlewares.PrometheusMiddleware)

	// Register all routes
	routes.RegisterRoutes(app)

	// Expose metrics endpoint
	app.Get("/metrics", func(c *fiber.Ctx) error {
		monitoring.ServeMetrics(c)
		return nil
	})

	log.Fatal(app.Listen(":2112"))

}
