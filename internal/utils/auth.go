package utils

import (
	"github.com/gin-gonic/gin"
)

// GetAuthUser retrieves the authenticated user from the context
func GetAuthUser(c *gin.Context) (AuthUser, bool) {
	authUserInterface, exists := c.Get("authUser")
	if !exists {
		return AuthUser{}, false
	}

	authUser, ok := authUserInterface.(AuthUser)
	if !ok {
		return AuthUser{}, false
	}

	return authUser, true
}
