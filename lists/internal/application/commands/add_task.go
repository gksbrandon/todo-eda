package commands

import (
	"context"

	"github.com/gksbrandon/todo-eda/internal/dispatcher"
	"github.com/gksbrandon/todo-eda/lists/internal/domain"
	"github.com/stackus/errors"
)

type (
	AddTask struct {
		ID          string
		ListID      string
		Description string
	}

	AddTaskHandler struct {
		lists           domain.ListRepository
		tasks           domain.TaskRepository
		domainPublisher dispatcher.EventPublisher
	}
)

func NewAddTaskHandler(lists domain.ListRepository, tasks domain.TaskRepository, domainPublisher dispatcher.EventPublisher) AddTaskHandler {
	return AddTaskHandler{
		lists:           lists,
		tasks:           tasks,
		domainPublisher: domainPublisher,
	}
}

func (h AddTaskHandler) AddTask(ctx context.Context, cmd AddTask) error {
	if _, err := h.lists.Find(ctx, cmd.ListID); err != nil {
		return errors.Wrap(err, "error adding task")
	}

	task, err := domain.CreateTask(cmd.ID, cmd.ListID, cmd.Description)
	if err != nil {
		return errors.Wrap(err, "error adding task")
	}

	if err = h.tasks.Save(ctx, task); err != nil {
		return errors.Wrap(err, "error adding task")
	}

	if err = h.domainPublisher.Publish(ctx, task.GetEvents()...); err != nil {
		return err
	}

	return nil
}
