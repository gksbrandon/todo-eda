package grpc

import (
	"context"

	"github.com/gksbrandon/todo-eda/users/internal/application"
	"github.com/gksbrandon/todo-eda/users/internal/domain"
	"github.com/gksbrandon/todo-eda/users/userspb"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	app application.App
	userspb.UnimplementedUsersServiceServer
}

var _ userspb.UsersServiceServer = (*server)(nil)

func RegisterServer(app application.App, registrar grpc.ServiceRegistrar) error {
	userspb.RegisterUsersServiceServer(registrar, server{app: app})
	return nil
}

func (s server) RegisterUser(ctx context.Context, request *userspb.RegisterUserRequest) (*userspb.RegisterUserResponse, error) {
	id := uuid.New().String()
	err := s.app.RegisterUser(ctx, application.RegisterUser{
		ID:    id,
		Name:  request.GetName(),
		Email: request.GetEmail(),
	})
	return &userspb.RegisterUserResponse{Id: id}, err
}

func (s server) AuthorizeUser(ctx context.Context, request *userspb.AuthorizeUserRequest) (*userspb.AuthorizeUserResponse, error) {
	err := s.app.AuthorizeUser(ctx, application.AuthorizeUser{
		ID: request.GetId(),
	})

	return &userspb.AuthorizeUserResponse{}, err
}

func (s server) GetUser(ctx context.Context, request *userspb.GetUserRequest) (*userspb.GetUserResponse, error) {
	user, err := s.app.GetUser(ctx, application.GetUser{
		ID: request.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &userspb.GetUserResponse{
		User: s.userFromDomain(user),
	}, nil
}

func (s server) userFromDomain(user *domain.User) *userspb.User {
	return &userspb.User{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
