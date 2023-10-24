// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: listspb/api.proto

package listspb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ListsServiceClient is the client API for ListsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ListsServiceClient interface {
	CreateList(ctx context.Context, in *CreateListRequest, opts ...grpc.CallOption) (*CreateListResponse, error)
	CompleteTask(ctx context.Context, in *CompleteTaskRequest, opts ...grpc.CallOption) (*CompleteTaskResponse, error)
	UncompleteTask(ctx context.Context, in *UncompleteTaskRequest, opts ...grpc.CallOption) (*UncompleteTaskResponse, error)
	AddTask(ctx context.Context, in *AddTaskRequest, opts ...grpc.CallOption) (*AddTaskResponse, error)
	RemoveTask(ctx context.Context, in *RemoveTaskRequest, opts ...grpc.CallOption) (*RemoveTaskResponse, error)
	GetTasks(ctx context.Context, in *GetTasksRequest, opts ...grpc.CallOption) (*GetTasksResponse, error)
}

type listsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewListsServiceClient(cc grpc.ClientConnInterface) ListsServiceClient {
	return &listsServiceClient{cc}
}

func (c *listsServiceClient) CreateList(ctx context.Context, in *CreateListRequest, opts ...grpc.CallOption) (*CreateListResponse, error) {
	out := new(CreateListResponse)
	err := c.cc.Invoke(ctx, "/listspb.ListsService/CreateList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listsServiceClient) CompleteTask(ctx context.Context, in *CompleteTaskRequest, opts ...grpc.CallOption) (*CompleteTaskResponse, error) {
	out := new(CompleteTaskResponse)
	err := c.cc.Invoke(ctx, "/listspb.ListsService/CompleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listsServiceClient) UncompleteTask(ctx context.Context, in *UncompleteTaskRequest, opts ...grpc.CallOption) (*UncompleteTaskResponse, error) {
	out := new(UncompleteTaskResponse)
	err := c.cc.Invoke(ctx, "/listspb.ListsService/UncompleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listsServiceClient) AddTask(ctx context.Context, in *AddTaskRequest, opts ...grpc.CallOption) (*AddTaskResponse, error) {
	out := new(AddTaskResponse)
	err := c.cc.Invoke(ctx, "/listspb.ListsService/AddTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listsServiceClient) RemoveTask(ctx context.Context, in *RemoveTaskRequest, opts ...grpc.CallOption) (*RemoveTaskResponse, error) {
	out := new(RemoveTaskResponse)
	err := c.cc.Invoke(ctx, "/listspb.ListsService/RemoveTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listsServiceClient) GetTasks(ctx context.Context, in *GetTasksRequest, opts ...grpc.CallOption) (*GetTasksResponse, error) {
	out := new(GetTasksResponse)
	err := c.cc.Invoke(ctx, "/listspb.ListsService/GetTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListsServiceServer is the server API for ListsService service.
// All implementations must embed UnimplementedListsServiceServer
// for forward compatibility
type ListsServiceServer interface {
	CreateList(context.Context, *CreateListRequest) (*CreateListResponse, error)
	CompleteTask(context.Context, *CompleteTaskRequest) (*CompleteTaskResponse, error)
	UncompleteTask(context.Context, *UncompleteTaskRequest) (*UncompleteTaskResponse, error)
	AddTask(context.Context, *AddTaskRequest) (*AddTaskResponse, error)
	RemoveTask(context.Context, *RemoveTaskRequest) (*RemoveTaskResponse, error)
	GetTasks(context.Context, *GetTasksRequest) (*GetTasksResponse, error)
	mustEmbedUnimplementedListsServiceServer()
}

// UnimplementedListsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedListsServiceServer struct {
}

func (UnimplementedListsServiceServer) CreateList(context.Context, *CreateListRequest) (*CreateListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateList not implemented")
}
func (UnimplementedListsServiceServer) CompleteTask(context.Context, *CompleteTaskRequest) (*CompleteTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteTask not implemented")
}
func (UnimplementedListsServiceServer) UncompleteTask(context.Context, *UncompleteTaskRequest) (*UncompleteTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UncompleteTask not implemented")
}
func (UnimplementedListsServiceServer) AddTask(context.Context, *AddTaskRequest) (*AddTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTask not implemented")
}
func (UnimplementedListsServiceServer) RemoveTask(context.Context, *RemoveTaskRequest) (*RemoveTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveTask not implemented")
}
func (UnimplementedListsServiceServer) GetTasks(context.Context, *GetTasksRequest) (*GetTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTasks not implemented")
}
func (UnimplementedListsServiceServer) mustEmbedUnimplementedListsServiceServer() {}

// UnsafeListsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ListsServiceServer will
// result in compilation errors.
type UnsafeListsServiceServer interface {
	mustEmbedUnimplementedListsServiceServer()
}

func RegisterListsServiceServer(s grpc.ServiceRegistrar, srv ListsServiceServer) {
	s.RegisterService(&ListsService_ServiceDesc, srv)
}

func _ListsService_CreateList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListsServiceServer).CreateList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listspb.ListsService/CreateList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListsServiceServer).CreateList(ctx, req.(*CreateListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListsService_CompleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListsServiceServer).CompleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listspb.ListsService/CompleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListsServiceServer).CompleteTask(ctx, req.(*CompleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListsService_UncompleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UncompleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListsServiceServer).UncompleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listspb.ListsService/UncompleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListsServiceServer).UncompleteTask(ctx, req.(*UncompleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListsService_AddTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListsServiceServer).AddTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listspb.ListsService/AddTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListsServiceServer).AddTask(ctx, req.(*AddTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListsService_RemoveTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListsServiceServer).RemoveTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listspb.ListsService/RemoveTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListsServiceServer).RemoveTask(ctx, req.(*RemoveTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListsService_GetTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListsServiceServer).GetTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listspb.ListsService/GetTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListsServiceServer).GetTasks(ctx, req.(*GetTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ListsService_ServiceDesc is the grpc.ServiceDesc for ListsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ListsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "listspb.ListsService",
	HandlerType: (*ListsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateList",
			Handler:    _ListsService_CreateList_Handler,
		},
		{
			MethodName: "CompleteTask",
			Handler:    _ListsService_CompleteTask_Handler,
		},
		{
			MethodName: "UncompleteTask",
			Handler:    _ListsService_UncompleteTask_Handler,
		},
		{
			MethodName: "AddTask",
			Handler:    _ListsService_AddTask_Handler,
		},
		{
			MethodName: "RemoveTask",
			Handler:    _ListsService_RemoveTask_Handler,
		},
		{
			MethodName: "GetTasks",
			Handler:    _ListsService_GetTasks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "listspb/api.proto",
}