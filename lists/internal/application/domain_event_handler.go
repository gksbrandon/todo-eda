package application

import (
	"context"

	"github.com/gksbrandon/todo-eda/internal/dispatcher"
)

type DomainEventHandlers interface {
	OnListCreated(ctx context.Context, event dispatcher.Event) error
	OnTaskAdded(ctx context.Context, event dispatcher.Event) error
	OnTaskCompleted(ctx context.Context, event dispatcher.Event) error
	OnTaskUncompleted(ctx context.Context, event dispatcher.Event) error
	OnTaskRemoved(ctx context.Context, event dispatcher.Event) error
}

type ignoreUnimplementedDomainEvents struct{}

var _ DomainEventHandlers = (*ignoreUnimplementedDomainEvents)(nil)

func (ignoreUnimplementedDomainEvents) OnListCreated(ctx context.Context, event dispatcher.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnTaskAdded(ctx context.Context, event dispatcher.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnTaskCompleted(ctx context.Context, event dispatcher.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnTaskUncompleted(ctx context.Context, event dispatcher.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnTaskRemoved(ctx context.Context, event dispatcher.Event) error {
	return nil
}
