package application

import (
	"context"

	"github.com/gksbrandon/todo-eda/internal/dispatcher"
	"github.com/gksbrandon/todo-eda/lists/internal/application/commands"
	"github.com/gksbrandon/todo-eda/lists/internal/application/queries"
	"github.com/gksbrandon/todo-eda/lists/internal/domain"
)

type (
	App interface {
		Commands
		Queries
	}
	Commands interface {
		CreateList(ctx context.Context, cmd commands.CreateList) error
		CompleteTask(ctx context.Context, cmd commands.CompleteTask) error
		UncompleteTask(ctx context.Context, cmd commands.UncompleteTask) error
		AddTask(ctx context.Context, cmd commands.AddTask) error
		RemoveTask(ctx context.Context, cmd commands.RemoveTask) error
	}
	Queries interface {
		GetTasks(ctx context.Context, query queries.GetTasks) ([]*domain.Task, error)
	}

	Application struct {
		appCommands
		appQueries
	}
	appCommands struct {
		commands.CreateListHandler
		commands.CompleteTaskHandler
		commands.UncompleteTaskHandler
		commands.AddTaskHandler
		commands.RemoveTaskHandler
	}
	appQueries struct {
		queries.GetTasksHandler
	}
)

var _ App = (*Application)(nil)

func New(lists domain.ListRepository, tasks domain.TaskRepository, domainPublisher dispatcher.EventPublisher,
) *Application {
	return &Application{
		appCommands: appCommands{
			CreateListHandler:     commands.NewCreateListHandler(lists, domainPublisher),
			CompleteTaskHandler:   commands.NewCompleteTaskHandler(tasks, domainPublisher),
			UncompleteTaskHandler: commands.NewUncompleteTaskHandler(tasks, domainPublisher),
			AddTaskHandler:        commands.NewAddTaskHandler(lists, tasks, domainPublisher),
			RemoveTaskHandler:     commands.NewRemoveTaskHandler(tasks, domainPublisher),
		},
		appQueries: appQueries{
			GetTasksHandler: queries.NewGetTasksHandler(tasks),
		},
	}
}
