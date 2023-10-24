package database

import "fmt"

type PGConfig struct {
	Host     string `envconfig:"PG_HOST" default:"postgres"`
	DbName   string `envconfig:"PG_DB_NAME" default:"todo"`
	User     string `envconfig:"PG_USER" default:"todo_user"`
	Password string `envconfig:"PG_PASSWORD" default:"todo_pass" conf:"mask"`
}

func (p PGConfig) GetConnString() string {
	return fmt.Sprintf("host=%s dbname=%s user=%s password=%s", p.Host, p.DbName, p.User, p.Password)
}
