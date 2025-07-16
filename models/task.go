package models

import "time"

type TaskStatus string

const (
	StatusNotStarted TaskStatus = "not-started"
	StatusInProgress TaskStatus = "in-progress"
	StatusCompleted  TaskStatus = "completed"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate"`
	Status      TaskStatus
}
