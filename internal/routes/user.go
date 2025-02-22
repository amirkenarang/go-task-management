package routes

import (
	"example.com/task-managment/internal/handlers"
	"github.com/gin-gonic/gin"
)

// UserRoutes registers user-related endpoints
func UserRoutes(router *gin.RouterGroup) {

	router.POST("/signup", handlers.SignUp)
	router.POST("/login", handlers.Login)
}
