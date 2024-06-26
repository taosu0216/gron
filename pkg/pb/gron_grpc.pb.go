// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: gron.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Gron_CreateTimer_FullMethodName = "/gron.Gron/CreateTimer"
	Gron_EnableTimer_FullMethodName = "/gron.Gron/EnableTimer"
)

// GronClient is the client API for Gron service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GronClient interface {
	CreateTimer(ctx context.Context, in *CreateTimerRequest, opts ...grpc.CallOption) (*CreateTimerResp, error)
	EnableTimer(ctx context.Context, in *EnableTimerRequest, opts ...grpc.CallOption) (*EnableTimerResp, error)
}

type gronClient struct {
	cc grpc.ClientConnInterface
}

func NewGronClient(cc grpc.ClientConnInterface) GronClient {
	return &gronClient{cc}
}

func (c *gronClient) CreateTimer(ctx context.Context, in *CreateTimerRequest, opts ...grpc.CallOption) (*CreateTimerResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTimerResp)
	err := c.cc.Invoke(ctx, Gron_CreateTimer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gronClient) EnableTimer(ctx context.Context, in *EnableTimerRequest, opts ...grpc.CallOption) (*EnableTimerResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EnableTimerResp)
	err := c.cc.Invoke(ctx, Gron_EnableTimer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GronServer is the server API for Gron service.
// All implementations must embed UnimplementedGronServer
// for forward compatibility
type GronServer interface {
	CreateTimer(context.Context, *CreateTimerRequest) (*CreateTimerResp, error)
	EnableTimer(context.Context, *EnableTimerRequest) (*EnableTimerResp, error)
	mustEmbedUnimplementedGronServer()
}

// UnimplementedGronServer must be embedded to have forward compatible implementations.
type UnimplementedGronServer struct {
}

func (UnimplementedGronServer) CreateTimer(context.Context, *CreateTimerRequest) (*CreateTimerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTimer not implemented")
}
func (UnimplementedGronServer) EnableTimer(context.Context, *EnableTimerRequest) (*EnableTimerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableTimer not implemented")
}
func (UnimplementedGronServer) mustEmbedUnimplementedGronServer() {}

// UnsafeGronServer may be embedded to opt out of forward compatibility for this service.
// Use of this interfaces is not recommended, as added methods to GronServer will
// result in compilation errors.
type UnsafeGronServer interface {
	mustEmbedUnimplementedGronServer()
}

func RegisterGronServer(s grpc.ServiceRegistrar, srv GronServer) {
	s.RegisterService(&Gron_ServiceDesc, srv)
}

func _Gron_CreateTimer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTimerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GronServer).CreateTimer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gron_CreateTimer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GronServer).CreateTimer(ctx, req.(*CreateTimerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gron_EnableTimer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnableTimerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GronServer).EnableTimer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gron_EnableTimer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GronServer).EnableTimer(ctx, req.(*EnableTimerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gron_ServiceDesc is the grpc.ServiceDesc for Gron service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gron_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gron.Gron",
	HandlerType: (*GronServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTimer",
			Handler:    _Gron_CreateTimer_Handler,
		},
		{
			MethodName: "EnableTimer",
			Handler:    _Gron_EnableTimer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gron.proto",
}
