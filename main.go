package main

import (
	"example.com/task-managment/db"
	"example.com/task-managment/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080

}
