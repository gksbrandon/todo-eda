package lists

import (
	"context"

	"github.com/gksbrandon/todo-eda/internal/dispatcher"
	"github.com/gksbrandon/todo-eda/internal/monolith"
	"github.com/gksbrandon/todo-eda/lists/internal/application"
	"github.com/gksbrandon/todo-eda/lists/internal/grpc"
	"github.com/gksbrandon/todo-eda/lists/internal/logging"
	"github.com/gksbrandon/todo-eda/lists/internal/postgres"
	"github.com/gksbrandon/todo-eda/lists/internal/rest"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	domainDispatcher := dispatcher.New()
	lists := postgres.NewListRepository("lists.lists", mono.Db())
	tasks := postgres.NewTaskRepository("lists.tasks", mono.Db())

	// setup application
	var app application.App
	app = application.New(lists, tasks, domainDispatcher)
	app = logging.LogApplicationAccess(app, mono.Logger())

	// setup Driver adapters
	if err := grpc.RegisterServer(ctx, app, mono.Rpc()); err != nil {
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
