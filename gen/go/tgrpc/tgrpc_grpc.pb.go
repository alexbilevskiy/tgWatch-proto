// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: tgrpc/tgrpc.proto

package tgrpc

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

const (
	Messages_GetScheduledMessages_FullMethodName       = "/Messages/GetScheduledMessages"
	Messages_ScheduleForwardedMessage_FullMethodName   = "/Messages/ScheduleForwardedMessage"
	Messages_EditMessageSchedulingState_FullMethodName = "/Messages/EditMessageSchedulingState"
)

// MessagesClient is the client API for Messages service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessagesClient interface {
	GetScheduledMessages(ctx context.Context, in *GetScheduledMessagesRequest, opts ...grpc.CallOption) (*GetScheduledMessagesResponse, error)
	ScheduleForwardedMessage(ctx context.Context, in *ScheduleForwardedMessageRequest, opts ...grpc.CallOption) (*ScheduleForwardedMessageResponse, error)
	EditMessageSchedulingState(ctx context.Context, in *EditMessageSchedulingStateRequest, opts ...grpc.CallOption) (*EditMessageSchedulingStateResponse, error)
}

type messagesClient struct {
	cc grpc.ClientConnInterface
}

func NewMessagesClient(cc grpc.ClientConnInterface) MessagesClient {
	return &messagesClient{cc}
}

func (c *messagesClient) GetScheduledMessages(ctx context.Context, in *GetScheduledMessagesRequest, opts ...grpc.CallOption) (*GetScheduledMessagesResponse, error) {
	out := new(GetScheduledMessagesResponse)
	err := c.cc.Invoke(ctx, Messages_GetScheduledMessages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagesClient) ScheduleForwardedMessage(ctx context.Context, in *ScheduleForwardedMessageRequest, opts ...grpc.CallOption) (*ScheduleForwardedMessageResponse, error) {
	out := new(ScheduleForwardedMessageResponse)
	err := c.cc.Invoke(ctx, Messages_ScheduleForwardedMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagesClient) EditMessageSchedulingState(ctx context.Context, in *EditMessageSchedulingStateRequest, opts ...grpc.CallOption) (*EditMessageSchedulingStateResponse, error) {
	out := new(EditMessageSchedulingStateResponse)
	err := c.cc.Invoke(ctx, Messages_EditMessageSchedulingState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessagesServer is the server API for Messages service.
// All implementations must embed UnimplementedMessagesServer
// for forward compatibility
type MessagesServer interface {
	GetScheduledMessages(context.Context, *GetScheduledMessagesRequest) (*GetScheduledMessagesResponse, error)
	ScheduleForwardedMessage(context.Context, *ScheduleForwardedMessageRequest) (*ScheduleForwardedMessageResponse, error)
	EditMessageSchedulingState(context.Context, *EditMessageSchedulingStateRequest) (*EditMessageSchedulingStateResponse, error)
	mustEmbedUnimplementedMessagesServer()
}

// UnimplementedMessagesServer must be embedded to have forward compatible implementations.
type UnimplementedMessagesServer struct {
}

func (UnimplementedMessagesServer) GetScheduledMessages(context.Context, *GetScheduledMessagesRequest) (*GetScheduledMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScheduledMessages not implemented")
}
func (UnimplementedMessagesServer) ScheduleForwardedMessage(context.Context, *ScheduleForwardedMessageRequest) (*ScheduleForwardedMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScheduleForwardedMessage not implemented")
}
func (UnimplementedMessagesServer) EditMessageSchedulingState(context.Context, *EditMessageSchedulingStateRequest) (*EditMessageSchedulingStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditMessageSchedulingState not implemented")
}
func (UnimplementedMessagesServer) mustEmbedUnimplementedMessagesServer() {}

// UnsafeMessagesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessagesServer will
// result in compilation errors.
type UnsafeMessagesServer interface {
	mustEmbedUnimplementedMessagesServer()
}

func RegisterMessagesServer(s grpc.ServiceRegistrar, srv MessagesServer) {
	s.RegisterService(&Messages_ServiceDesc, srv)
}

func _Messages_GetScheduledMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScheduledMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServer).GetScheduledMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messages_GetScheduledMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServer).GetScheduledMessages(ctx, req.(*GetScheduledMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messages_ScheduleForwardedMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleForwardedMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServer).ScheduleForwardedMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messages_ScheduleForwardedMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServer).ScheduleForwardedMessage(ctx, req.(*ScheduleForwardedMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messages_EditMessageSchedulingState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditMessageSchedulingStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServer).EditMessageSchedulingState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messages_EditMessageSchedulingState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServer).EditMessageSchedulingState(ctx, req.(*EditMessageSchedulingStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Messages_ServiceDesc is the grpc.ServiceDesc for Messages service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Messages_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Messages",
	HandlerType: (*MessagesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetScheduledMessages",
			Handler:    _Messages_GetScheduledMessages_Handler,
		},
		{
			MethodName: "ScheduleForwardedMessage",
			Handler:    _Messages_ScheduleForwardedMessage_Handler,
		},
		{
			MethodName: "EditMessageSchedulingState",
			Handler:    _Messages_EditMessageSchedulingState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tgrpc/tgrpc.proto",
}
