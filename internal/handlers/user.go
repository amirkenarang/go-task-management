package handlers

import (
	"net/http"

	"example.com/task-managment/internal/models"
	"example.com/task-managment/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func SignUp(context *fiber.Ctx) error {

	var user models.User
	err := context.BodyParser(&user)

	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Could not parse request data.", "error": err.Error()})
	}

	err = user.Save()
	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not save user. ", "error": err.Error()})
	}

	return context.Status(http.StatusCreated).JSON(fiber.Map{"message": "User created successfully.", "user": user})
}

func Login(context *fiber.Ctx) error {
	var user models.User
	err := context.BodyParser(&user)
	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Could not parse request data.", "error": err.Error()})
	}

	err = user.ValidateCredentioals()
	if err != nil {
		return context.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "Could not authenticate user", "error": err.Error()})
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not generate token.", "error": err.Error()})
	}

	return context.Status(http.StatusCreated).JSON(fiber.Map{"message": "Login successfully.", "user": user, "token": token})

}
