package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gksbrandon/todo-eda/internal/dispatcher"
	"github.com/gksbrandon/todo-eda/users/internal/domain"
)

type UserRepository struct {
	tableName string
	db        *sql.DB
}

var _ domain.UserRepository = (*UserRepository)(nil)

func NewUserRepository(tableName string, db *sql.DB) UserRepository {
	return UserRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r UserRepository) Save(ctx context.Context, user *domain.User) error {
	const query = "INSERT INTO %s (id, name, email) VALUES ($1, $2, $3)"

	_, err := r.db.ExecContext(ctx, r.table(query), user.ID, user.Name, user.Email)

	return err
}

func (r UserRepository) Find(ctx context.Context, userID string) (*domain.User, error) {
	const query = "SELECT name, email FROM %s WHERE id = $1 LIMIT 1"

	user := &domain.User{
		AggregateBase: dispatcher.AggregateBase{
			ID: userID,
		},
	}

	err := r.db.QueryRowContext(ctx, r.table(query), userID).Scan(&user.Name, &user.Email)

	return user, err
}

func (r UserRepository) Update(ctx context.Context, user *domain.User) error {
	const query = "UPDATE %s SET name = $2, email = $3 WHERE id = $1"

	_, err := r.db.ExecContext(ctx, r.table(query), user.ID, user.Name, user.Email)

	return err
}

func (r UserRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
