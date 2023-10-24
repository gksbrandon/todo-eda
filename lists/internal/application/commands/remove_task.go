package commands

import (
	"context"

	"github.com/gksbrandon/todo-eda/internal/dispatcher"
	"github.com/gksbrandon/todo-eda/lists/internal/domain"
)

type (
	RemoveTask struct {
		ID string
	}

	RemoveTaskHandler struct {
		tasks           domain.TaskRepository
		domainPublisher dispatcher.EventPublisher
	}
)

func NewRemoveTaskHandler(tasks domain.TaskRepository, domainPublisher dispatcher.EventPublisher) RemoveTaskHandler {
	return RemoveTaskHandler{
		tasks:           tasks,
		domainPublisher: domainPublisher,
	}
}

func (h RemoveTaskHandler) RemoveTask(ctx context.Context, cmd RemoveTask) error {
	task, err := h.tasks.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = task.Remove(); err != nil {
		return err
	}

	if err = h.tasks.Delete(ctx, cmd.ID); err != nil {
		return err
	}

	if err = h.domainPublisher.Publish(ctx, task.GetEvents()...); err != nil {
		return err
	}

	return nil
}
