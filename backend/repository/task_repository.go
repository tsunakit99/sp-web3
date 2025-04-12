package repository

import "github.com/tsunakit99/sp-web3/domain"

type TaskRepository interface {
	FindByUserID(userID string) ([]domain.Task, error)
}
