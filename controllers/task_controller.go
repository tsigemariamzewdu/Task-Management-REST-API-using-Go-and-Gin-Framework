package controllers

import (
	"net/http"
	"strconv"
	"task_management/data"
	"task_management/models"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context){
	tasks:= data.GetTasks()
	c.IndentedJSON(http.StatusOK,tasks)
}

func GetTaskByID(c *gin.Context){
	idp:=c.Param("id")

	id,err:= strconv.Atoi(idp)
	if err !=nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"error":"Invalid Id"})
		return
	}
	task,found:=data.GetTaskByID(id)
	if !found{
		c.IndentedJSON(http.StatusNotFound,gin.H{"error":"task not found"})
		return 
	}
	c.IndentedJSON(http.StatusOK,task)
}

func AddTask(c *gin.Context){
	var newTask models.Task

	if err:= c.ShouldBindJSON(&newTask);err !=nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	ok := data.AddTask(newTask)
	if !ok {
		c.IndentedJSON(http.StatusConflict,gin.H{"error":"task already exits"})
		return
	}
	c.IndentedJSON(http.StatusCreated,newTask)
}

func DeleteTaskByID(c *gin.Context){
	idp :=c.Param("id")
	id,err:= strconv.Atoi(idp)

	if err!=nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"error":"invalid ID"})
		return
	}
	deleted :=data.DeleteTaskByID(id)
	if !deleted{
		c.IndentedJSON(http.StatusNotFound,gin.H{"error":"task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK,gin.H{"message":"task deleted"})

}
func UpdateTaskByID(c *gin.Context){
	idp:= c.Param("id")
	id,err:=strconv.Atoi(idp)
	if err!=nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"error":"Invalid ID"})
		return
	}

	var updatedTask models.Task
	if err:= c.ShouldBindJSON(&updatedTask);err!=nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	
	ok:=data.UpdateTaskByID(id,updatedTask)
	if !ok{
		c.IndentedJSON(http.StatusNotFound,gin.H{"error":"task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK,updatedTask)

}