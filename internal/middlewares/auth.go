package middlewares

import (
	"net/http"

	"example.com/task-managment/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func Authenticate(context *fiber.Ctx) error {
	token := context.Get("Authorization")
	if token == "" {
		return context.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "No token provided.",
		})
	}

	authUser, err := utils.VerifyToken(token)
	if err != nil {
		return context.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token. Try again later.",
		})
	}

	context.Locals("authUser", authUser)
	return context.Next()
}
