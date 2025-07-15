package router

import (
	
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine){
	router.GET("/tasks")
	router.GET("/tasks/:id")
	router.PUT("/tasks/:id")
	router.DELETE("/tasks/:id")
	router.POST("/tasks")
}
