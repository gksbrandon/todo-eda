package domain

import (
	"github.com/gksbrandon/todo-eda/internal/dispatcher"
	"github.com/stackus/errors"
)

var (
	ErrTaskDescriptionIsBlank   = errors.Wrap(errors.ErrBadRequest, "the task description cannot be blank")
	ErrTaskIsAlreadyCompleted   = errors.Wrap(errors.ErrBadRequest, "the task is already completed")
	ErrTaskIsAlreadyUncompleted = errors.Wrap(errors.ErrBadRequest, "the task is already uncompleted")
)

type Task struct {
	dispatcher.AggregateBase
	ListID      string
	Description string
	Completed   bool
}

func CreateTask(id, listID, description string) (*Task, error) {
	if description == "" {
		return nil, ErrTaskDescriptionIsBlank
	}

	task := &Task{
		AggregateBase: dispatcher.AggregateBase{
			ID: id,
		},
		ListID:      listID,
		Description: description,
	}

	task.AddEvent(&TaskAdded{
		Task: task,
	})

	return task, nil
}

func (s *Task) CompleteTask() (err error) {
	if s.Completed {
		return ErrTaskIsAlreadyCompleted
	}

	s.Completed = true

	s.AddEvent(&TaskCompleted{
		Task: s,
	})

	return
}

func (s *Task) UncompleteTask() (err error) {
	if !s.Completed {
		return ErrTaskIsAlreadyUncompleted
	}

	s.Completed = false

	s.AddEvent(&TaskUncompleted{
		Task: s,
	})

	return
}

func (p *Task) Remove() error {
	p.AddEvent(&TaskRemoved{
		Task: p,
	})

	return nil
}
