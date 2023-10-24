package logging

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/gksbrandon/todo-eda/users/internal/application"
	"github.com/gksbrandon/todo-eda/users/internal/domain"
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

func (a Application) RegisterUser(ctx context.Context, register application.RegisterUser) (err error) {
	a.logger.Info().Msg("--> Users.RegisterUser")
	defer func() { a.logger.Info().Err(err).Msg("<-- Users.RegisterUser") }()
	return a.App.RegisterUser(ctx, register)
}

func (a Application) AuthorizeUser(ctx context.Context, authorize application.AuthorizeUser) (err error) {
	a.logger.Info().Msg("--> Users.AuthorizeUser")
	defer func() { a.logger.Info().Err(err).Msg("<-- Users.AuthorizeUser") }()
	return a.App.AuthorizeUser(ctx, authorize)
}

func (a Application) GetUser(ctx context.Context, get application.GetUser) (user *domain.User,
	err error,
) {
	a.logger.Info().Msg("--> Users.GetUser")
	defer func() { a.logger.Info().Err(err).Msg("<-- Users.GetUser") }()
	return a.App.GetUser(ctx, get)
}
