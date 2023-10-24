package domain

import (
	"github.com/gksbrandon/todo-eda/internal/dispatcher"
)

type List struct {
	dispatcher.AggregateBase
	UserID string
}

func CreateList(id, userID string) (list *List, err error) {
	list = &List{
		AggregateBase: dispatcher.AggregateBase{
			ID: id,
		},
		UserID: userID,
	}

	list.AddEvent(&ListCreated{
		List: list,
	})

	return
}
