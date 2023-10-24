package commands

import (
	"context"

	"github.com/gksbrandon/todo-eda/internal/dispatcher"
	"github.com/gksbrandon/todo-eda/lists/internal/domain"
)

type (
	CompleteTask struct {
		ID string
	}

	CompleteTaskHandler struct {
		tasks           domain.TaskRepository
		domainPublisher dispatcher.EventPublisher
	}
)

func NewCompleteTaskHandler(tasks domain.TaskRepository, domainPublisher dispatcher.EventPublisher) CompleteTaskHandler {
	return CompleteTaskHandler{
		tasks:           tasks,
		domainPublisher: domainPublisher,
	}
}

func (h CompleteTaskHandler) CompleteTask(ctx context.Context, cmd CompleteTask) error {
	task, err := h.tasks.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = task.CompleteTask(); err != nil {
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
