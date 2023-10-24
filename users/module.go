package users

import (
	"context"

	"github.com/gksbrandon/todo-eda/internal/dispatcher"
	"github.com/gksbrandon/todo-eda/internal/monolith"
	"github.com/gksbrandon/todo-eda/users/internal/application"
	"github.com/gksbrandon/todo-eda/users/internal/grpc"
	"github.com/gksbrandon/todo-eda/users/internal/logging"
	"github.com/gksbrandon/todo-eda/users/internal/postgres"
	"github.com/gksbrandon/todo-eda/users/internal/rest"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	domainDispatcher := dispatcher.New()
	users := postgres.NewUserRepository("users.users", mono.Db())

	// setup application
	app := logging.LogApplicationAccess(
		application.New(users, domainDispatcher),
		mono.Logger(),
	)

	if err := grpc.RegisterServer(app, mono.Rpc()); err != nil {
		return err
	}
	if err := rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}

	return nil
}
