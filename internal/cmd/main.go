package main

import (
	"example.com/task-managment/internal/db"
	"example.com/task-managment/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080

}
