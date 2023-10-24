package queries

import (
	"context"

	"github.com/gksbrandon/todo-eda/lists/internal/domain"
)

type GetTasks struct {
	ListID string
}

type GetTasksHandler struct {
	tasks domain.TaskRepository
}

func NewGetTasksHandler(tasks domain.TaskRepository) GetTasksHandler {
	return GetTasksHandler{tasks: tasks}
}

func (h GetTasksHandler) GetTasks(ctx context.Context, query GetTasks) ([]*domain.Task, error) {
	return h.tasks.GetTasks(ctx, query.ListID)
}
