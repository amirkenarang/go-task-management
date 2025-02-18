package routes

import (
	"log"
	"net/http"
	"strconv"

	"example.com/task-managment/models"
	"github.com/gin-gonic/gin"
)

func getTask(context *gin.Context) {
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

func getTasks(context *gin.Context) {
	tasks, err := models.GetAllTasks()
	log.Printf("Error fetching tasks: %v", err) // Log the error

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch tasks. Try again later."})
		return
	}
	context.JSON(http.StatusOK, tasks)
}

func createTasks(context *gin.Context) {
	var task models.Task
	err := context.ShouldBindJSON(&task)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	task.ID = 1
	task.UserID = 1

	err = task.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create task. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Task created!", "task": task})
}

func updateTask(context *gin.Context) {
	taskId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch task. Try again later."})
		return
	}

	_, err = models.GetTaskById(taskId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the task."})
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
