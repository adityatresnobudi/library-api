// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: library.proto

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

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Auth_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "library.proto",
}

// BorrowClient is the client API for Borrow service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BorrowClient interface {
	Borrow(ctx context.Context, in *BorrowRequest, opts ...grpc.CallOption) (*BorrowResponse, error)
	Return(ctx context.Context, in *ReturnRequest, opts ...grpc.CallOption) (*ReturnResponse, error)
}

type borrowClient struct {
	cc grpc.ClientConnInterface
}

func NewBorrowClient(cc grpc.ClientConnInterface) BorrowClient {
	return &borrowClient{cc}
}

func (c *borrowClient) Borrow(ctx context.Context, in *BorrowRequest, opts ...grpc.CallOption) (*BorrowResponse, error) {
	out := new(BorrowResponse)
	err := c.cc.Invoke(ctx, "/auth.Borrow/Borrow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *borrowClient) Return(ctx context.Context, in *ReturnRequest, opts ...grpc.CallOption) (*ReturnResponse, error) {
	out := new(ReturnResponse)
	err := c.cc.Invoke(ctx, "/auth.Borrow/Return", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BorrowServer is the server API for Borrow service.
// All implementations must embed UnimplementedBorrowServer
// for forward compatibility
type BorrowServer interface {
	Borrow(context.Context, *BorrowRequest) (*BorrowResponse, error)
	Return(context.Context, *ReturnRequest) (*ReturnResponse, error)
	mustEmbedUnimplementedBorrowServer()
}

// UnimplementedBorrowServer must be embedded to have forward compatible implementations.
type UnimplementedBorrowServer struct {
}

func (UnimplementedBorrowServer) Borrow(context.Context, *BorrowRequest) (*BorrowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Borrow not implemented")
}
func (UnimplementedBorrowServer) Return(context.Context, *ReturnRequest) (*ReturnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Return not implemented")
}
func (UnimplementedBorrowServer) mustEmbedUnimplementedBorrowServer() {}

// UnsafeBorrowServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BorrowServer will
// result in compilation errors.
type UnsafeBorrowServer interface {
	mustEmbedUnimplementedBorrowServer()
}

func RegisterBorrowServer(s grpc.ServiceRegistrar, srv BorrowServer) {
	s.RegisterService(&Borrow_ServiceDesc, srv)
}

func _Borrow_Borrow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BorrowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BorrowServer).Borrow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Borrow/Borrow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BorrowServer).Borrow(ctx, req.(*BorrowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Borrow_Return_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReturnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BorrowServer).Return(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Borrow/Return",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BorrowServer).Return(ctx, req.(*ReturnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Borrow_ServiceDesc is the grpc.ServiceDesc for Borrow service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Borrow_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Borrow",
	HandlerType: (*BorrowServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Borrow",
			Handler:    _Borrow_Borrow_Handler,
		},
		{
			MethodName: "Return",
			Handler:    _Borrow_Return_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "library.proto",
}
