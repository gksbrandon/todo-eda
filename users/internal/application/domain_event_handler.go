package application

import (
	"context"

	"github.com/gksbrandon/todo-eda/internal/dispatcher"
)

type DomainEventHandlers interface {
	OnUserRegistered(ctx context.Context, event dispatcher.Event) error
	OnUserAuthorized(ctx context.Context, event dispatcher.Event) error
}

type ignoreUnimplementedDomainEvents struct{}

var _ DomainEventHandlers = (*ignoreUnimplementedDomainEvents)(nil)

func (ignoreUnimplementedDomainEvents) OnUserRegistered(ctx context.Context, event dispatcher.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnUserAuthorized(ctx context.Context, event dispatcher.Event) error {
	return nil
}
