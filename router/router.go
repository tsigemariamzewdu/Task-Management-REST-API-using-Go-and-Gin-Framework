package router

import (
	"task_management/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine){
	router.GET("/tasks",controllers.GetTasks)
	router.GET("/tasks/:id",controllers.GetTaskByID)
	router.PUT("/tasks/:id",controllers.UpdateTaskByID)
	router.DELETE("/tasks/:id",controllers.GetTaskByID)
	router.POST("/tasks",controllers.AddTask)
}
