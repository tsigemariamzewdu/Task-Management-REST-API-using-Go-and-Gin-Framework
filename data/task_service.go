package data

import "task_management/models"

var tasks []models.Task

func GetTasks() []models.Task{ 
	return tasks

}

func GetTaskByID(id int)(models.Task,bool){
	for _,task := range tasks{
		if task.ID ==id{
			return task,true
		}

	}
	return models.Task{},false
	
}

func AddTask(task models.Task) bool {
	
	for _, t := range tasks {
		if t.Title == task.Title && t.Description == task.Description {
			return false 
		}
	}

	var maxID int
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	task.ID = maxID + 1

	tasks = append(tasks, task)
	return true
}

func DeleteTaskByID(id int)bool{
	for i,task := range tasks{
		if task.ID==id{
			tasks = append(tasks[:i],tasks[i+1:]... )
			return true
		}

		}
	return false
	}

func UpdateTaskByID(id int,updatedTask models.Task)bool{
	for i,task := range tasks{
		if task.ID==id{
			updatedTask.ID=id
			tasks[i] = updatedTask
			return true

		}
	}
	return false
}