package commands

import (
	"context"

	"github.com/gksbrandon/todo-eda/internal/dispatcher"
	"github.com/gksbrandon/todo-eda/lists/internal/domain"
)

type (
	CreateList struct {
		ID     string
		UserID string
	}

	CreateListHandler struct {
		lists           domain.ListRepository
		domainPublisher dispatcher.EventPublisher
	}
)

func NewCreateListHandler(lists domain.ListRepository, domainPublisher dispatcher.EventPublisher) CreateListHandler {
	return CreateListHandler{
		lists:           lists,
		domainPublisher: domainPublisher,
	}
}

func (h CreateListHandler) CreateList(ctx context.Context, cmd CreateList) error {
	list, err := domain.CreateList(cmd.ID, cmd.UserID)
	if err != nil {
		return err
	}

	if err = h.lists.Save(ctx, list); err != nil {
		return err
	}

	if err = h.domainPublisher.Publish(ctx, list.GetEvents()...); err != nil {
		return err
	}

	return nil
}
