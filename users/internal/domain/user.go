package domain

import (
	"github.com/gksbrandon/todo-eda/internal/dispatcher"

	"github.com/stackus/errors"
)

type User struct {
	dispatcher.AggregateBase
	Name  string
	Email string
}

var (
	ErrNameCannotBeBlank   = errors.Wrap(errors.ErrBadRequest, "the user name cannot be blank")
	ErrUserIDCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "the user id cannot be blank")
	ErrEmailCannotBeBlank  = errors.Wrap(errors.ErrBadRequest, "the email cannot be blank")
	ErrUserNotAuthorized   = errors.Wrap(errors.ErrUnauthorized, "user is not authorized")
)

func RegisterUser(id, name, email string) (*User, error) {
	if id == "" {
		return nil, ErrUserIDCannotBeBlank
	}

	if name == "" {
		return nil, ErrNameCannotBeBlank
	}

	if email == "" {
		return nil, ErrEmailCannotBeBlank
	}

	user := &User{
		AggregateBase: dispatcher.AggregateBase{
			ID: id,
		},
		Name:  name,
		Email: email,
	}

	user.AddEvent(&UserRegistered{
		User: user,
	})

	return user, nil
}

func (c *User) Authorize(token string) error {
	// TODO: Authorization

	c.AddEvent(&UserAuthorized{
		User: c,
	})

	return nil
}
