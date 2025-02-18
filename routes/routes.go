package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	server.GET("/tasks", getTasks)
	server.GET("/tasks/:id", getTask)
	server.POST(("/tasks"), createTasks)
	server.PUT(("/tasks/:id"), updateTask)
}
