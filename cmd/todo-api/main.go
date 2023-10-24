package main

import (
	"database/sql"
	stdLog "log"
	"net/http"
	"os"

	"github.com/ardanlabs/conf"
	"github.com/gksbrandon/todo-eda/internal/config"
	"github.com/gksbrandon/todo-eda/internal/logger"
	"github.com/gksbrandon/todo-eda/internal/monolith"
	"github.com/gksbrandon/todo-eda/internal/rpc"
	"github.com/gksbrandon/todo-eda/internal/waiter"
	"github.com/gksbrandon/todo-eda/internal/web"
	"github.com/gksbrandon/todo-eda/lists"
	"github.com/gksbrandon/todo-eda/users"
	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Initialiaze Configuration
	cfg, err := config.InitConfig()
	if err != nil {
		stdLog.Fatalf("config.InitConfig: %+v", err)
	}

	// Initialize logger
	logLevel, err := logger.NewLevel(cfg.LogLevel)
	if err != nil {
		stdLog.Fatalf("logger.NewLevel: %+v", err)
	}
	log := logger.New(logger.LogConfig{
		Environment: cfg.Env,
		LogLevel:    logLevel,
	}, cfg.Env)

	// Run application
	if err := run(log, cfg); err != nil {
		log.Error().Msgf("run: %w", err)
		os.Exit(1)
	}
}

func run(log zerolog.Logger, cfg config.AppConfig) (err error) {
	// Setup new application
	a := app{
		cfg:    cfg,
		logger: log,
	}

	// Log Configuration
	confOutput, err := conf.String(&a.cfg)
	if err != nil {
		return
	}
	log.Info().Msg(confOutput)

	// Setup Postgres Connection
	a.db, err = sql.Open("pgx", cfg.PG.GetConnString())
	if err != nil {
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(a.db)

	// Initialize drivers
	a.mux = initMux(cfg.Web)
	a.rpc = initRpc(cfg.Rpc)

	// Initialize modules
	a.waiter = waiter.New(waiter.CatchSignals())
	a.modules = []monolith.Module{
		&lists.Module{},
		&users.Module{},
	}
	if err = a.startupModules(); err != nil {
		return err
	}

	// Mount general web resources
	a.mux.Mount("/", http.FileServer(http.FS(web.WebUI)))
	log.Info().Msg("started todo application")
	defer log.Info().Msg("stopped todo application")

	// Add drivers and initiate server
	a.waiter.Add(
		a.waitForRpc,
		a.waitForWeb,
	)

	return a.waiter.Wait()
}

func initMux(_ web.WebConfig) *chi.Mux {
	return chi.NewMux()
}

func initRpc(_ rpc.RpcConfig) *grpc.Server {
	server := grpc.NewServer()
	reflection.Register(server)

	return server
}
