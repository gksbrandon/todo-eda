package logging

import (
	"context"

	"github.com/gksbrandon/todo-eda/lists/internal/application"
	"github.com/gksbrandon/todo-eda/lists/internal/application/commands"
	"github.com/gksbrandon/todo-eda/lists/internal/application/queries"
	"github.com/gksbrandon/todo-eda/lists/internal/domain"
	"github.com/rs/zerolog"
)

type Application struct {
	application.App
	logger zerolog.Logger
}

var _ application.App = (*Application)(nil)

func LogApplicationAccess(application application.App, logger zerolog.Logger) Application {
	return Application{
		App:    application,
		logger: logger,
	}
}

func (a Application) CreateList(ctx context.Context, cmd commands.CreateList) (err error) {
	a.logger.Info().Msg("--> Lists.CreateList")
	defer func() { a.logger.Info().Err(err).Msg("<-- Lists.CreateList") }()
	return a.App.CreateList(ctx, cmd)
}

func (a Application) AddTask(ctx context.Context, cmd commands.AddTask) (err error) {
	a.logger.Info().Msg("--> Lists.AddTask")
	defer func() { a.logger.Info().Err(err).Msg("<-- Lists.AddTask") }()
	return a.App.AddTask(ctx, cmd)
}

func (a Application) CompleteTask(ctx context.Context, cmd commands.CompleteTask) (err error) {
	a.logger.Info().Msg("--> Lists.CompleteTask")
	defer func() { a.logger.Info().Err(err).Msg("<-- Lists.CompleteTask") }()
	return a.App.CompleteTask(ctx, cmd)
}

func (a Application) UncompleteTask(ctx context.Context, cmd commands.UncompleteTask) (err error) {
	a.logger.Info().Msg("--> Lists.UncompleteTask")
	defer func() { a.logger.Info().Err(err).Msg("<-- Lists.UncompleteTask") }()
	return a.App.UncompleteTask(ctx, cmd)
}

func (a Application) RemoveTask(ctx context.Context, cmd commands.RemoveTask) (err error) {
	a.logger.Info().Msg("--> Lists.RemoveTask")
	defer func() { a.logger.Info().Err(err).Msg("<-- Lists.RemoveTask") }()
	return a.App.RemoveTask(ctx, cmd)
}

func (a Application) GetTasks(ctx context.Context, query queries.GetTasks) (tasks []*domain.Task, err error) {
	a.logger.Info().Msg("--> Lists.GetTasks")
	defer func() { a.logger.Info().Err(err).Msg("<-- Lists.GetTasks") }()
	return a.App.GetTasks(ctx, query)
}
