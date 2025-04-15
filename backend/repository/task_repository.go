package repository

import "github.com/tsunakit99/sp-web3/domain"

type TaskRepository interface {
	FindByUserID(userID string) ([]domain.Task, error)
	Insert(task domain.Task) error
	Update(task domain.Task) error
	// Delete(id string) error
	// ToggleCompleted(id string, userID string) error
}
