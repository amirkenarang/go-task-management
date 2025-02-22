package middlewares

import (
	"net/http"

	"example.com/task-managment/internal/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.GetHeader("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No token provided."})
		return
	}

	authUser, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token. Try again later."})
		return
	}

	context.Set("authUser", authUser)
	context.Next()
}
