package domain

import (
	"context"
)

type UserRepository interface {
	Save(ctx context.Context, user *User) error
	Find(ctx context.Context, userID string) (*User, error)
	Update(ctx context.Context, user *User) error
}
