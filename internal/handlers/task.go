package handlers

import (
	"net/http"
	"strconv"

	"example.com/task-managment/internal/utils"

	"example.com/task-managment/internal/models"
	"github.com/gofiber/fiber/v2"
)

func GetTask(context *fiber.Ctx) error {
	taskId, err := strconv.ParseInt(context.Params("id"), 10, 64)

	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not fetch task. Try again later."})

	}

	task, err := models.GetTaskById(taskId)

	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not fetch tasks. Try again later."})

	}
	return context.Status(http.StatusOK).JSON(task)
}

func GetTasks(context *fiber.Ctx) error {
	tasks, err := models.GetAllTasks()

	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not fetch tasks. Try again later."})

	}
	return context.Status(http.StatusOK).JSON(tasks)
}

func CreateTasks(context *fiber.Ctx) error {
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

	err = task.Save()

	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not create task. Try again later."})
	}

	return context.Status(http.StatusCreated).JSON(fiber.Map{"message": "Task created!", "task": task})
}

func UpdateTask(context *fiber.Ctx) error {
	taskId, err := strconv.ParseInt(context.Params("id"), 10, 64)

	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not fetch task. Try again later."})
	}

	authUser, exists := utils.GetAuthUser(context)
	if !exists {
		return context.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "User does not exist in context"})
	}

	task, err := models.GetTaskById(taskId)

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

	err = updatedTask.Update()

	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not update task. Try again later."})
	}

	return context.Status(http.StatusOK).JSON(fiber.Map{"message": "Updated Successfully"})

}

func DeleteTask(context *fiber.Ctx) error {
	taskId, err := strconv.ParseInt(context.Params("id"), 10, 64)

	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not fetch id. Try again later."})
	}

	task, err := models.GetTaskById(taskId)
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

	err = task.Delete()
	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not delete task. Try again later."})
	}

	return context.Status(http.StatusOK).JSON(fiber.Map{"message": "Deleted Successfully"})

}
