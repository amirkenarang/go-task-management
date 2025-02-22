package utils

import (
	"github.com/gofiber/fiber/v2"
)

// GetAuthUser retrieves the authenticated user from the context
func GetAuthUser(c *fiber.Ctx) (AuthUser, bool) {
	authUserInterface := c.Locals("authUser")
	if authUserInterface == nil {
		return AuthUser{}, false
	}

	authUser, ok := authUserInterface.(AuthUser)
	if !ok {
		return AuthUser{}, false
	}

	return authUser, true
}
