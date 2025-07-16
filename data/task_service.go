package data

import "task_management/models"

var Tasks []models.Task

func GetTasks() []models.Task{ 
	return Tasks

}

func GetTaskByID(id int)(models.Task,bool){
	for _,task := range Tasks{
		if task.ID ==id{
			return task,true
		}

	}
	return models.Task{},false
	
}

func AddTask(task models.Task) bool {
	
	for _, t := range Tasks {
		if t.Title == task.Title && t.Description == task.Description {
			return false 
		}
	}

	var maxID int
	for _, t := range Tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	task.ID = maxID + 1

	Tasks = append(Tasks, task)
	return true
}

func DeleteTaskByID(id int)bool{
	for i,task := range Tasks{
		if task.ID==id{
			Tasks = append(Tasks[:i],Tasks[i+1:]... )
			return true
		}

		}
	return false
	}

func UpdateTaskByID(id int,updatedTask models.Task)bool{
	for i,task := range Tasks{
		if task.ID==id{
			updatedTask.ID=id
			Tasks[i] = updatedTask
			return true

		}
	}
	return false
}