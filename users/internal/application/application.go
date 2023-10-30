package application

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gksbrandon/todo-eda/internal/dispatcher"
	"github.com/gksbrandon/todo-eda/users/internal/domain"
)

type (
	RegisterUser struct {
		ID    string
		Name  string
		Email string
	}

	AuthorizeUser struct {
		ID    string
		Token string
	}

	GetUser struct {
		ID string
	}

	App interface {
		RegisterUser(ctx context.Context, register RegisterUser) error
		AuthorizeUser(ctx context.Context, authorize AuthorizeUser) error
		GetUser(ctx context.Context, get GetUser) (*domain.User, error)
	}

	Application struct {
		users           domain.UserRepository
		domainPublisher dispatcher.EventPublisher
	}
)

var _ App = (*Application)(nil)

func New(users domain.UserRepository, domainPublisher dispatcher.EventPublisher) *Application {
	return &Application{
		users:           users,
		domainPublisher: domainPublisher,
	}
}

func (a Application) RegisterUser(ctx context.Context, register RegisterUser) error {
	user, err := domain.RegisterUser(register.ID, register.Name, register.Email)
	if err != nil {
		return err
	}

	if err = a.users.Save(ctx, user); err != nil {
		return err
	}

	if err = a.domainPublisher.Publish(ctx, user.GetEvents()...); err != nil {
		return err
	}

	return nil
}

func (a Application) AuthorizeUser(ctx context.Context, authorize AuthorizeUser) error {
	// user, err := a.users.Find(ctx, authorize.ID)
	// if err != nil {
	// 	return err
	// }

	// if err = user.Authorize(authorize.Token); err != nil {
	// 	return err
	// }

	// if err := a.domainPublisher.Publish(ctx, user.GetEvents()...); err != nil {
	// 	return err
	// }

	// Extract the Bearer token from the Authorization header
	// tokenString := r.Header.Get("Authorization")
	// if tokenString == "" {
	// 	http.Error(w, "Token missing", http.StatusUnauthorized)
	// 	return
	// }

	// Use Google APIs to get user data
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + authorize.Token)
	if err != nil {
		return err
	}
	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(userData)

	return nil
}

func (a Application) GetUser(ctx context.Context, get GetUser) (*domain.User, error) {
	return a.users.Find(ctx, get.ID)
}
