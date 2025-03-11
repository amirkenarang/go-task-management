package middlewares

import (
	"strings"
	"time"

	"example.com/task-management/internal/monitoring"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
)

func PrometheusMiddleware(c *fiber.Ctx) error {
	start := time.Now()
	method := c.Method()
	path := normalizePath(c.Path()) // Normalize the path

	// Skip invalid or non-API paths
	if !isValidAPIPath(path) {
		return c.Next()
	}

	monitoring.RequestCounter.With(prometheus.Labels{"method": method, "endpoint": path}).Inc()

	err := c.Next()

	duration := time.Since(start).Seconds()
	monitoring.RequestDuration.With(prometheus.Labels{"method": method, "endpoint": path}).Observe(duration)

	return err
}

// normalizePath normalizes the path for consistent labeling
func normalizePath(path string) string {
	if strings.HasPrefix(path, "/api/login/") {
		return "/api/login"
	}
	if strings.HasPrefix(path, "/api/tasks/") {
		return "/api/tasks"
	}
	// Add more rules as needed
	return path
}

// isValidAPIPath checks if the path is a valid API route
func isValidAPIPath(path string) bool {
	// Define valid API routes
	validRoutes := []string{
		"/api/login",
		"/api/tasks",
		"/metrics", // Include the metrics endpoint if needed
	}

	for _, route := range validRoutes {
		if path == route || strings.HasPrefix(path, route+"/") {
			return true
		}
	}
	return false
}
