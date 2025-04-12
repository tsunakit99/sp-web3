package task

type TaskUsecase interface {
	GetTasks(userID string) ([]TaskDTO, error)
}