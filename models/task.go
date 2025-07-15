package models

import "time"

type Task struct {
	ID          int    `json:"Id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     time.Time `json:"dueDate"`
	
}