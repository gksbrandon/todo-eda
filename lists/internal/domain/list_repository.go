package domain

import (
	"context"
)

type ListRepository interface {
	Save(ctx context.Context, list *List) error
	Find(ctx context.Context, listID string) (*List, error)
}
