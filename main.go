package main

import (
	"task_management/router"
	
	"github.com/gin-gonic/gin"
)

func main() {
	app:=gin.Default()
	router.SetUpRoutes(app)
	app.Run("localhost:8080")
	
}