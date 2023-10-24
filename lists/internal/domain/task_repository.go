package domain

import (
	"context"
)

type TaskRepository interface {
	Find(ctx context.Context, id string) (*Task, error)
	Save(ctx context.Context, task *Task) error
	Delete(ctx context.Context, id string) error
	GetTasks(ctx context.Context, listID string) ([]*Task, error)
	Update(ctx context.Context, task *Task) error
}
