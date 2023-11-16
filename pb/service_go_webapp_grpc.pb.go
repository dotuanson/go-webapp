// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: service_go_webapp.proto

package pb

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
	GoWebApp_CreateUser_FullMethodName = "/pb.GoWebApp/CreateUser"
	GoWebApp_LoginUser_FullMethodName  = "/pb.GoWebApp/LoginUser"
)

// GoWebAppClient is the client API for GoWebApp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoWebAppClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
}

type goWebAppClient struct {
	cc grpc.ClientConnInterface
}

func NewGoWebAppClient(cc grpc.ClientConnInterface) GoWebAppClient {
	return &goWebAppClient{cc}
}

func (c *goWebAppClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, GoWebApp_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goWebAppClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, GoWebApp_LoginUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoWebAppServer is the server API for GoWebApp service.
// All implementations must embed UnimplementedGoWebAppServer
// for forward compatibility
type GoWebAppServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	mustEmbedUnimplementedGoWebAppServer()
}

// UnimplementedGoWebAppServer must be embedded to have forward compatible implementations.
type UnimplementedGoWebAppServer struct {
}

func (UnimplementedGoWebAppServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedGoWebAppServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedGoWebAppServer) mustEmbedUnimplementedGoWebAppServer() {}

// UnsafeGoWebAppServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoWebAppServer will
// result in compilation errors.
type UnsafeGoWebAppServer interface {
	mustEmbedUnimplementedGoWebAppServer()
}

func RegisterGoWebAppServer(s grpc.ServiceRegistrar, srv GoWebAppServer) {
	s.RegisterService(&GoWebApp_ServiceDesc, srv)
}

func _GoWebApp_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoWebAppServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoWebApp_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoWebAppServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoWebApp_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoWebAppServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoWebApp_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoWebAppServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GoWebApp_ServiceDesc is the grpc.ServiceDesc for GoWebApp service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoWebApp_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.GoWebApp",
	HandlerType: (*GoWebAppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _GoWebApp_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _GoWebApp_LoginUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_go_webapp.proto",
}