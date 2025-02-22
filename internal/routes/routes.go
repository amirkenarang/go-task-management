package routes

import (
	"example.com/task-managment/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api") // Group all API routes under `/api`

	authenticated := router.Group("/api") // Group all API routes under `/api`
	authenticated.Use(middlewares.Authenticate)

	TaskRoutes(authenticated)
	UserRoutes(api)
}
