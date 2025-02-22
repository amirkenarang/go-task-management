package handlers

import (
	"log"
	"net/http"
	"strconv"

	"example.com/task-managment/internal/utils"

	"example.com/task-managment/internal/models"
	"github.com/gin-gonic/gin"
)

func GetTask(context *gin.Context) {
	taskId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch task. Try again later."})
		return
	}

	task, err := models.GetTaskById(taskId)

	log.Printf("Error fetching task: %v", err) // Log the error

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch tasks. Try again later."})
		return
	}
	context.JSON(http.StatusOK, task)
}

func GetTasks(context *gin.Context) {
	tasks, err := models.GetAllTasks()
	log.Printf("Error fetching tasks: %v", err) // Log the error

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch tasks. Try again later."})
		return
	}
	context.JSON(http.StatusOK, tasks)
}

func CreateTasks(context *gin.Context) {
	var task models.Task
	err := context.ShouldBindJSON(&task)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	authUser, exists := utils.GetAuthUser(context)
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User does not exist in context"})
		return
	}

	task.UserID = authUser.UserId

	err = task.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create task. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Task created!", "task": task})
}

func UpdateTask(context *gin.Context) {
	taskId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch task. Try again later."})
		return
	}

	authUser, exists := utils.GetAuthUser(context)
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User does not exist in context"})
		return
	}

	task, err := models.GetTaskById(taskId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the task."})
		return
	}

	if task.UserID != authUser.UserId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User not Not authorized for update this task!"})
		return
	}

	var updatedTask models.Task

	err = context.ShouldBindJSON(&updatedTask)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	updatedTask.ID = taskId

	err = updatedTask.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update task. Try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Updated Successfully"})

}

func DeleteTask(context *gin.Context) {
	taskId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch id. Try again later."})
		return
	}

	task, err := models.GetTaskById(taskId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch tasks. Try again later."})
		return
	}

	authUser, exists := utils.GetAuthUser(context)
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User does not exist in context"})
		return
	}
	if task.UserID != authUser.UserId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User not Not authorized to delete this task"})
		return
	}

	err = task.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete task. Try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Deleted Successfully"})

}
