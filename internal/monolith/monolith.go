package monolith

import (
	"context"
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/gksbrandon/todo-eda/internal/config"
	"github.com/gksbrandon/todo-eda/internal/waiter"
	"google.golang.org/grpc"
)

type Monolith interface {
	Config() config.AppConfig
	Db() *sql.DB
	Logger() zerolog.Logger
	Mux() *chi.Mux
	Rpc() *grpc.Server
	Waiter() waiter.Waiter
}

type Module interface {
	Startup(context.Context, Monolith) error
}
