package routes

import (
	"example.com/task-managment/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.GET("/tasks", getTasks)
	authenticated.GET("/tasks/:id", getTask)
	authenticated.POST("/tasks", createTasks)
	authenticated.PUT("/tasks/:id", updateTask)
	authenticated.DELETE("/tasks/:id", deleteTask)

	server.POST("/signup", signUp)
	server.POST("/login", login)
}
