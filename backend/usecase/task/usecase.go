package task

import (
	"time"

	"github.com/tsunakit99/sp-web3/domain"
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
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			IsCompleted: t.IsCompleted,
		})
	}
	return dtos, nil
}

func (u *taskUsecase) CreateTask(userID string, input CreateTaskInput) error {

	if err := ValidateTaskAndDescription(input.Title, input.Description); err != nil {
		return err
	}

	task := domain.Task{
		UserID:      userID,
		Title:       input.Title,
		Description: input.Description,
		IsCompleted: false,
		DueDate:     input.DueDate,
		CreatedAt:   time.Now(),
	}
	return u.repo.Insert(task)
}

func (u *taskUsecase) UpdateTask(userID string, input UpdateTaskInput) error {

	if err := ValidateTaskAndDescription(input.Title, input.Description); err != nil {
		return err
	}

	task := domain.Task{
		ID:          input.ID,
		UserID:      userID,
		Title:       input.Title,
		Description: input.Description,
		DueDate:     input.DueDate,
	}
	return u.repo.Update(task)
}

func (u *taskUsecase) DeleteTask(userID string, taskID string) error {
	return u.repo.Delete(userID, taskID)
}
