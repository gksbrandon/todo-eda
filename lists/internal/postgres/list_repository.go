package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gksbrandon/todo-eda/internal/dispatcher"
	"github.com/gksbrandon/todo-eda/lists/internal/domain"
	"github.com/stackus/errors"
)

type ListRepository struct {
	tableName string
	db        *sql.DB
}

var _ domain.ListRepository = (*ListRepository)(nil)

func NewListRepository(tableName string, db *sql.DB) ListRepository {
	return ListRepository{tableName: tableName, db: db}
}

func (r ListRepository) Find(ctx context.Context, listID string) (*domain.List, error) {
	const query = "SELECT user_id FROM %s WHERE id = $1 LIMIT 1"

	list := &domain.List{
		AggregateBase: dispatcher.AggregateBase{
			ID: listID,
		},
	}

	err := r.db.QueryRowContext(ctx, r.table(query), listID).Scan(&list.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "scanning list")
	}

	return list, nil
}

func (r ListRepository) FindByUserID(ctx context.Context, userID string) (*domain.List, error) {
	const query = "SELECT id FROM %s WHERE user_id = $1 LIMIT 1"

	list := &domain.List{
		UserID: userID,
	}

	err := r.db.QueryRowContext(ctx, r.table(query), userID).Scan(&list.ID)
	if err != nil {
		return nil, errors.Wrap(err, "scanning list")
	}

	return list, nil
}

func (r ListRepository) Save(ctx context.Context, list *domain.List) error {
	const query = "INSERT INTO %s (id, user_id) VALUES ($1, $2)"

	_, err := r.db.ExecContext(ctx, r.table(query), list.ID, list.UserID)

	return errors.Wrap(err, "inserting list")
}

func (r ListRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
