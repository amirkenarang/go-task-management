package handlers

import (
	"net/http"
	"strconv"

	"example.com/task-managment/internal/repository"
	"example.com/task-managment/internal/utils"

	"example.com/task-managment/internal/models"
	"github.com/gofiber/fiber/v2"
)

type TaskHandler struct {
	Repo *repository.TaskRepository
}

func NewTaskHandler(repo *repository.TaskRepository) *TaskHandler {
	return &TaskHandler{Repo: repo}
}

// CreateTasks handles the creation of a new task.
// It parses the request body into a Task model, retrieves the authenticated user,
// assigns the user ID to the task, and then save it to the database.
//
// If the request body cannot be parsed, it returns a 400 Bad Request response.
// If the user is not authenticated, it returns a 401 Unauthorized response.
// If there is an error saving the task, it returns a 500 Internal Server Error response.
// On success, it returns a 201 Created response with the created task.
//
// @param context *fiber.Ctx - Fiber request context.
// @return error - Fiber response containing JSON data.
func (h *TaskHandler) CreateTasks(context *fiber.Ctx) error {
	var task models.Task
	err := context.BodyParser(&task)

	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Could not parse request data."})
	}

	authUser, exists := utils.GetAuthUser(context)
	if !exists {
		return context.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "User does not exist in context"})
	}

	task.UserID = authUser.UserId

	err = h.Repo.Save(&task)

	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not create task. Try again later."})
	}

	return context.Status(http.StatusCreated).JSON(fiber.Map{"message": "Task created!", "task": task})
}

// GetTask retrieves a task by its ID from the database.
// It extracts the task ID from the request parameters, fetches the task using the repository,
// and returns it as a JSON response.
//
// If the task ID is invalid, it returns a 500 Internal Server Error response.
// If the task cannot be fetched, it returns a 500 Internal Server Error response.
// On success, it returns a 200 OK response with the task details.
func (h *TaskHandler) GetTask(context *fiber.Ctx) error {
	taskId, err := strconv.ParseInt(context.Params("id"), 10, 64)

	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not fetch task. Try again later."})

	}

	task, err := h.Repo.GetTaskById(taskId)

	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not fetch tasks. Try again later."})

	}
	return context.Status(http.StatusOK).JSON(task)
}

// GetTasks retrieves all tasks from the database and returns them as a JSON response.
//
// If there is an error fetching tasks, it returns a 500 Internal Server Error response.
// On success, it returns a 200 OK response with the list of tasks.
func (h *TaskHandler) GetTasks(context *fiber.Ctx) error {
	tasks, err := h.Repo.GetAllTasks()

	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not fetch tasks. Try again later."})

	}
	return context.Status(http.StatusOK).JSON(tasks)
}

// UpdateTask updates an existing task based on the request data.
// It validates the task ID, checks if the authenticated user is authorized to update the task,
// and then updates it in the database.
//
// If the task ID is invalid, it returns a 500 Internal Server Error response.
// If the user is not authenticated, it returns a 401 Unauthorized response.
// If the task does not exist, it returns a 500 Internal Server Error response.
// If the user is not authorized to update the task, it returns a 401 Unauthorized response.
// If the request body cannot be parsed, it returns a 400 Bad Request response.
// If updating the task fails, it returns a 500 Internal Server Error response.
// On success, it returns a 200 OK response with a success message.
func (h *TaskHandler) UpdateTask(context *fiber.Ctx) error {
	taskId, err := strconv.ParseInt(context.Params("id"), 10, 64)

	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not fetch task. Try again later."})
	}

	authUser, exists := utils.GetAuthUser(context)
	if !exists {
		return context.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "User does not exist in context"})
	}

	task, err := h.Repo.GetTaskById(taskId)

	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not fetch the task."})
	}

	if task.UserID != authUser.UserId {
		return context.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "User not Not authorized for update this task!"})
	}

	var updatedTask models.Task

	err = context.BodyParser(&updatedTask)

	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Could not parse request data."})
	}

	updatedTask.ID = taskId

	err = h.Repo.Update(&updatedTask)

	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not update task. Try again later."})
	}

	return context.Status(http.StatusOK).JSON(fiber.Map{"message": "Updated Successfully"})

}

// DeleteTask deletes a task by its ID after validating user authorization.
// It ensures that the task exists and that the authenticated user is allowed to delete it.
//
// If the task ID is invalid, it returns a 500 Internal Server Error response.
// If the task does not exist, it returns a 500 Internal Server Error response.
// If the user is not authenticated, it returns a 401 Unauthorized response.
// If the user is not authorized to delete the task, it returns a 401 Unauthorized response.
// If deleting the task fails, it returns a 500 Internal Server Error response.
// On success, it returns a 200 OK response with a success message.
func (h *TaskHandler) DeleteTask(context *fiber.Ctx) error {
	taskId, err := strconv.ParseInt(context.Params("id"), 10, 64)

	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not fetch id. Try again later."})
	}

	task, err := h.Repo.GetTaskById(taskId)
	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not fetch tasks. Try again later."})
	}

	authUser, exists := utils.GetAuthUser(context)
	if !exists {
		return context.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "User does not exist in context"})
	}
	if task.UserID != authUser.UserId {
		return context.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "User not Not authorized to delete this task"})
	}

	err = h.Repo.Delete(taskId)
	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not delete task. Try again later."})
	}

	return context.Status(http.StatusOK).JSON(fiber.Map{"message": "Deleted Successfully"})

}
