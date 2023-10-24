package grpc

import (
	"context"

	"github.com/gksbrandon/todo-eda/lists/internal/application"
	"github.com/gksbrandon/todo-eda/lists/internal/application/commands"
	"github.com/gksbrandon/todo-eda/lists/internal/application/queries"
	"github.com/gksbrandon/todo-eda/lists/internal/domain"
	"github.com/gksbrandon/todo-eda/lists/listspb"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	app application.App
	listspb.UnimplementedListsServiceServer
}

var _ listspb.ListsServiceServer = (*server)(nil)

func RegisterServer(_ context.Context, app application.App, registrar grpc.ServiceRegistrar) error {
	listspb.RegisterListsServiceServer(registrar, server{app: app})
	return nil
}

func (s server) CreateList(ctx context.Context, request *listspb.CreateListRequest) (*listspb.CreateListResponse, error) {
	listID := uuid.New().String()

	// id := uuid.New().String()
	err := s.app.CreateList(ctx, commands.CreateList{
		ID:     listID,
		UserID: request.UserId,
	})
	if err != nil {
		return nil, err
	}

	return &listspb.CreateListResponse{
		Id: listID,
	}, nil
}

func (s server) AddTask(ctx context.Context, request *listspb.AddTaskRequest) (*listspb.AddTaskResponse, error) {
	id := uuid.New().String()
	err := s.app.AddTask(ctx, commands.AddTask{
		ID:          id,
		ListID:      request.GetListId(),
		Description: request.GetDescription(),
	})
	if err != nil {
		return nil, err
	}

	return &listspb.AddTaskResponse{Id: id}, nil
}

func (s server) CompleteTask(ctx context.Context, request *listspb.CompleteTaskRequest) (*listspb.CompleteTaskResponse, error) {
	err := s.app.CompleteTask(ctx, commands.CompleteTask{
		ID: request.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &listspb.CompleteTaskResponse{}, nil
}

func (s server) UncompleteTask(ctx context.Context, request *listspb.UncompleteTaskRequest) (*listspb.UncompleteTaskResponse, error) {
	err := s.app.UncompleteTask(ctx, commands.UncompleteTask{
		ID: request.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &listspb.UncompleteTaskResponse{}, nil
}

func (s server) RemoveTask(ctx context.Context, request *listspb.RemoveTaskRequest) (*listspb.RemoveTaskResponse, error) {
	err := s.app.RemoveTask(ctx, commands.RemoveTask{
		ID: request.GetId(),
	})

	return &listspb.RemoveTaskResponse{}, err
}

func (s server) GetTasks(ctx context.Context, request *listspb.GetTasksRequest) (*listspb.GetTasksResponse, error) {
	tasks, err := s.app.GetTasks(ctx, queries.GetTasks{ListID: request.GetListId()})
	if err != nil {
		return nil, err
	}

	protoTasks := []*listspb.Task{}
	for _, task := range tasks {
		protoTasks = append(protoTasks, s.taskFromDomain(task))
	}

	return &listspb.GetTasksResponse{
		Tasks: protoTasks,
	}, nil
}

func (s server) listFromDomain(list *domain.List) *listspb.List {
	return &listspb.List{
		Id: list.ID,
	}
}

func (s server) taskFromDomain(task *domain.Task) *listspb.Task {
	return &listspb.Task{
		Id:          task.ID,
		ListId:      task.ListID,
		Description: task.Description,
		Completed:   task.Completed,
	}
}
