package routes

import (
	"example.com/task-managment/internal/handlers"
	"github.com/gin-gonic/gin"
)

func TaskRoutes(router *gin.RouterGroup) {

	router.GET("/tasks", handlers.GetTasks)
	router.GET("/tasks/:id", handlers.GetTask)
	router.POST("/tasks", handlers.CreateTasks)
	router.PUT("/tasks/:id", handlers.UpdateTask)
	router.DELETE("/tasks/:id", handlers.DeleteTask)
}
