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

func (r *taskRepository) Insert(task domain.Task) error {
	_, err := r.db.Exec("INSERT INTO tasks (user_id, title, description, is_completed, due_date, created_at) VALUES ($1, $2, $3, $4, $5, $6)",
		task.UserID, task.Title, task.Description, task.IsCompleted, task.DueDate, task.CreatedAt)
	return err
}

func (r *taskRepository) Update(task domain.Task) error {
	_, err := r.db.Exec("UPDATE tasks SET title = $1, description = $2, due_date = $3 WHERE id = $4 AND user_id = $5",
		task.Title, task.Description, task.DueDate, task.ID, task.UserID)
	return err
}

func (r *taskRepository) Delete(taskID string, userID string) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id = $1 AND user_id = $2", taskID, userID)
	return err
}

func (r *taskRepository) ToggleComplete(taskID string, userID string) error {
	_, err := r.db.Exec("UPDATE tasks SET is_completed = NOT is_completed WHERE id = $1 AND user_id = $2", taskID, userID)
	return err
}
