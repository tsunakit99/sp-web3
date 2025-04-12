package task

import (
	"github.com/tsunakit99/sp-web3/repository"
)

type taskUsecase struct {
	repo repository.TaskRepository
}

type TaskDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}

func New(r repository.TaskRepository) TaskUsecase {
	return &taskUsecase{r}
}

func (u *taskUsecase) GetTasks(userID string) ([]TaskDTO, error) {
	tasks, err := u.repo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	var dtos []TaskDTO
	for _, t := range tasks {
		dtos = append(dtos, TaskDTO{
			ID: t.ID,
			Title: t.Title,
			Description: t.Description,
			IsCompleted: t.IsCompleted,
		})
	}
	return dtos, nil
}
