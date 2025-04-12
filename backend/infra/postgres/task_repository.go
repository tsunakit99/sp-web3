package postgres

import (
	"database/sql"

	"github.com/tsunakit99/sp-web3/domain"
	"github.com/tsunakit99/sp-web3/repository"
)

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) repository.TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) FindByUserID(userID string) ([]domain.Task, error) {
	rows, err := r.db.Query("SELECT id, title, description, is_completed FROM tasks WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []domain.Task{}
	for rows.Next() {
		var t domain.Task
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.IsCompleted)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}
