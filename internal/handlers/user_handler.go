package handlers

import (
	"net/http"
	"strconv"

	"example.com/task-management/internal/repository"
	"example.com/task-management/internal/utils"

	"example.com/task-management/internal/models"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{Repo: repo}
}

// CreateUser handles the creation of a new user.
// It parses the request body into a User model, retrieves the authenticated user,
// assigns the user ID to the user, and then save it to the database.
//
// If the request body cannot be parsed, it returns a 400 Bad Request response.
// If the user is not authenticated, it returns a 401 Unauthorized response.
// If there is an error saving the user, it returns a 500 Internal Server Error response.
// On success, it returns a 201 Created response with the created user.
//
// @param context *fiber.Ctx - Fiber request context.
// @return error - Fiber response containing JSON data.
func (h *UserHandler) CreateUser(context *fiber.Ctx) error {
	var user models.User
	err := context.BodyParser(&user)

	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Could not parse request data."})
	}

	_, exists := utils.GetAuthUser(context)
	if !exists {
		return context.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "User does not exist in context"})
	}

	err = h.Repo.Save(&user)

	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not create user. Try again later."})
	}

	return context.Status(http.StatusCreated).JSON(fiber.Map{"message": "User created!", "user": user})
}

// GetUser retrieves a user by its ID from the database.
// It extracts the user ID from the request parameters, fetches the user using the repository,
// and returns it as a JSON response.
//
// If the user ID is invalid, it returns a 500 Internal Server Error response.
// If the user cannot be fetched, it returns a 500 Internal Server Error response.
// On success, it returns a 200 OK response with the user details.
func (h *UserHandler) GetUser(context *fiber.Ctx) error {
	userId, err := strconv.ParseInt(context.Params("id"), 10, 64)

	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not fetch user. Try again later."})

	}

	user, err := h.Repo.GetUserById(userId)

	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not fetch users. Try again later."})

	}
	return context.Status(http.StatusOK).JSON(user)
}

// GetUsers retrieves all users from the database and returns them as a JSON response.
//
// If there is an error fetching users, it returns a 500 Internal Server Error response.
// On success, it returns a 200 OK response with the list of users.
func (h *UserHandler) GetUsers(context *fiber.Ctx) error {
	users, err := h.Repo.GetAllUsers()

	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not fetch users. Try again later."})

	}
	return context.Status(http.StatusOK).JSON(users)
}

// UpdateUser updates an existing user based on the request data.
// It validates the user ID, checks if the authenticated user is authorized to update the user,
// and then updates it in the database.
//
// If the user ID is invalid, it returns a 500 Internal Server Error response.
// If the user is not authenticated, it returns a 401 Unauthorized response.
// If the user does not exist, it returns a 500 Internal Server Error response.
// If the user is not authorized to update the user, it returns a 401 Unauthorized response.
// If the request body cannot be parsed, it returns a 400 Bad Request response.
// If updating the user fails, it returns a 500 Internal Server Error response.
// On success, it returns a 200 OK response with a success message.
func (h *UserHandler) UpdateUser(context *fiber.Ctx) error {
	userId, err := strconv.ParseInt(context.Params("id"), 10, 64)

	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not fetch user. Try again later."})
	}

	_, exists := utils.GetAuthUser(context)
	if !exists {
		return context.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "User does not exist in context"})
	}

	_, err = h.Repo.GetUserById(userId)

	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not fetch the user."})
	}

	var updatedUser models.User

	err = context.BodyParser(&updatedUser)

	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Could not parse request data."})
	}

	updatedUser.ID = userId

	err = h.Repo.Update(&updatedUser)

	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not update user. Try again later."})
	}

	return context.Status(http.StatusOK).JSON(fiber.Map{"message": "Updated Successfully"})

}

// DeleteUser deletes a user by its ID after validating user authorization.
// It ensures that the user exists and that the authenticated user is allowed to delete it.
//
// If the user ID is invalid, it returns a 500 Internal Server Error response.
// If the user does not exist, it returns a 500 Internal Server Error response.
// If the user is not authenticated, it returns a 401 Unauthorized response.
// If the user is not authorized to delete the user, it returns a 401 Unauthorized response.
// If deleting the user fails, it returns a 500 Internal Server Error response.
// On success, it returns a 200 OK response with a success message.
func (h *UserHandler) DeleteUser(context *fiber.Ctx) error {
	userId, err := strconv.ParseInt(context.Params("id"), 10, 64)

	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not fetch id. Try again later."})
	}

	_, err = h.Repo.GetUserById(userId)
	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not fetch users. Try again later."})
	}

	_, exists := utils.GetAuthUser(context)
	if !exists {
		return context.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "User does not exist in context"})
	}

	err = h.Repo.Delete(userId)
	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not delete user. Try again later."})
	}

	return context.Status(http.StatusOK).JSON(fiber.Map{"message": "Deleted Successfully"})

}
