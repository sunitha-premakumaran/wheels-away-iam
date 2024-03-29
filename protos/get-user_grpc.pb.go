// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: protos/get-user.proto

package wheels_away_iam

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

// UserProcessorClient is the client API for UserProcessor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserProcessorClient interface {
	GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userProcessorClient struct {
	cc grpc.ClientConnInterface
}

func NewUserProcessorClient(cc grpc.ClientConnInterface) UserProcessorClient {
	return &userProcessorClient{cc}
}

func (c *userProcessorClient) GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.UserProcessor/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserProcessorServer is the server API for UserProcessor service.
// All implementations must embed UnimplementedUserProcessorServer
// for forward compatibility
type UserProcessorServer interface {
	GetUser(context.Context, *UserRequest) (*UserResponse, error)
	mustEmbedUnimplementedUserProcessorServer()
}

// UnimplementedUserProcessorServer must be embedded to have forward compatible implementations.
type UnimplementedUserProcessorServer struct {
}

func (UnimplementedUserProcessorServer) GetUser(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserProcessorServer) mustEmbedUnimplementedUserProcessorServer() {}

// UnsafeUserProcessorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserProcessorServer will
// result in compilation errors.
type UnsafeUserProcessorServer interface {
	mustEmbedUnimplementedUserProcessorServer()
}

func RegisterUserProcessorServer(s grpc.ServiceRegistrar, srv UserProcessorServer) {
	s.RegisterService(&UserProcessor_ServiceDesc, srv)
}

func _UserProcessor_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProcessorServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserProcessor/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProcessorServer).GetUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserProcessor_ServiceDesc is the grpc.ServiceDesc for UserProcessor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserProcessor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserProcessor",
	HandlerType: (*UserProcessorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserProcessor_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/get-user.proto",
}
