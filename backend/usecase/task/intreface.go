package task

import "time"

type CreateTaskInput struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"due_date"`
}

type UpdateTaskInput struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"due_date"`
}
type TaskUsecase interface {
	GetTasks(userID string) ([]TaskDTO, error)
	CreateTask(userID string, input CreateTaskInput) error
	UpdateTask(userID string, input UpdateTaskInput) error
	DeleteTask(userID string, taskID string) error
	ToggleComplete(userID string, taskID string) error
}
