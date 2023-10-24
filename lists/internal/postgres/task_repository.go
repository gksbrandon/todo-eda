package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gksbrandon/todo-eda/internal/dispatcher"
	"github.com/gksbrandon/todo-eda/lists/internal/domain"
	"github.com/stackus/errors"
)

type TaskRepository struct {
	tableName string
	db        *sql.DB
}

var _ domain.TaskRepository = (*TaskRepository)(nil)

func NewTaskRepository(tableName string, db *sql.DB) TaskRepository {
	return TaskRepository{tableName: tableName, db: db}
}

func (r TaskRepository) Find(ctx context.Context, id string) (*domain.Task, error) {
	const query = "SELECT list_id, description, completed FROM %s WHERE id = $1 LIMIT 1"

	task := &domain.Task{
		AggregateBase: dispatcher.AggregateBase{ID: id},
	}

	err := r.db.QueryRowContext(ctx, r.table(query), id).Scan(&task.ListID, &task.Description, &task.Completed)
	if err != nil {
		return nil, errors.Wrap(err, "scanning task")
	}

	return task, nil
}

func (r TaskRepository) Save(ctx context.Context, task *domain.Task) error {
	const query = "INSERT INTO %s (id, list_id, description, completed) VALUES ($1, $2, $3, $4)"

	_, err := r.db.ExecContext(ctx, r.table(query), task.ID, task.ListID, task.Description, task.Completed)

	return errors.Wrap(err, "inserting task")
}

func (r TaskRepository) Delete(ctx context.Context, id string) error {
	const query = "DELETE FROM %s WHERE id = $1"

	_, err := r.db.ExecContext(ctx, r.table(query), id)

	return errors.Wrap(err, "deleting task")
}

func (r TaskRepository) GetTasks(ctx context.Context, listID string) ([]*domain.Task, error) {
	const query = "SELECT id, description, completed FROM %s WHERE list_id = $1"

	tasks := make([]*domain.Task, 0)

	rows, err := r.db.QueryContext(ctx, r.table(query), listID)
	if err != nil {
		return nil, errors.Wrap(err, "querying tasks")
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			err = errors.Wrap(err, "closing task rows")
		}
	}(rows)

	for rows.Next() {
		task := &domain.Task{
			ListID: listID,
		}
		err := rows.Scan(&task.ID, &task.Description, &task.Completed)
		if err != nil {
			return nil, errors.Wrap(err, "scanning task")
		}

		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "finishing task rows")
	}

	return tasks, nil
}

func (r TaskRepository) Update(ctx context.Context, task *domain.Task) error {
	const query = "UPDATE %s SET description = $2, completed = $3 WHERE id = $1"

	_, err := r.db.ExecContext(ctx, r.table(query), task.ID, task.Description, task.Completed)

	return errors.Wrap(err, "updating task")
}

func (r TaskRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
