package config

import (
	"fmt"
	"os"
	"time"

	"github.com/gksbrandon/todo-eda/internal/database"
	"github.com/gksbrandon/todo-eda/internal/rpc"
	"github.com/gksbrandon/todo-eda/internal/web"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type (
	AppConfig struct {
		Env             string
		LogLevel        string `envconfig:"LOG_LEVEL" default:"DEBUG"`
		PG              database.PGConfig
		Rpc             rpc.RpcConfig
		Web             web.WebConfig
		ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"30s"`
	}
)

func InitConfig() (cfg AppConfig, err error) {
	cfg.Env = os.Getenv("env")
	if cfg.Env == "" {
		cfg.Env = "dev"
	}

	if err = validateEnvironment(cfg.Env); err != nil {
		return
	}

	configs := fmt.Sprintf("environments/%s/configs/configs.yaml", cfg.Env)
	secrets := fmt.Sprintf("environments/%s/configs/secrets.yaml", cfg.Env)
	if err = godotenv.Load(configs, secrets); err != nil {
		return
	}

	if err = envconfig.Process("", &cfg); err != nil {
		return
	}
	return
}
