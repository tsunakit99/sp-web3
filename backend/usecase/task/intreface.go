package task

import "time"

type CreateTaskInput struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"due_date"`
}
type TaskUsecase interface {
	GetTasks(userID string) ([]TaskDTO, error)
	CreateTask(userID string, input CreateTaskInput) error
}
