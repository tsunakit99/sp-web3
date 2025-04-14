package domain

import "time"

type Task struct {
	ID          string
	UserID      string
	Title       string
	Description string
	IsCompleted bool
	DueDate     *time.Time
	CreatedAt   time.Time
}
