package handlers

import (
	"net/http"

	"example.com/task-managment/internal/models"
	"example.com/task-managment/internal/utils"
	"github.com/gin-gonic/gin"
)

func SignUp(context *gin.Context) {

	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user. ", "error": err.Error()})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully.", "user": user})
}

func Login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}

	err = user.ValidateCredentioals()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user", "error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token.", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Login successfully.", "user": user, "token": token})

}
