package commands

import (
	"context"

	"github.com/gksbrandon/todo-eda/internal/dispatcher"
	"github.com/gksbrandon/todo-eda/lists/internal/domain"
)

type (
	UncompleteTask struct {
		ID string
	}

	UncompleteTaskHandler struct {
		tasks           domain.TaskRepository
		domainPublisher dispatcher.EventPublisher
	}
)

func NewUncompleteTaskHandler(tasks domain.TaskRepository, domainPublisher dispatcher.EventPublisher) UncompleteTaskHandler {
	return UncompleteTaskHandler{
		tasks:           tasks,
		domainPublisher: domainPublisher,
	}
}

func (h UncompleteTaskHandler) UncompleteTask(ctx context.Context, cmd UncompleteTask) error {
	task, err := h.tasks.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = task.UncompleteTask(); err != nil {
		return err
	}

	if err = h.tasks.Update(ctx, task); err != nil {
		return err
	}

	if err = h.domainPublisher.Publish(ctx, task.GetEvents()...); err != nil {
		return err
	}

	return nil
}
