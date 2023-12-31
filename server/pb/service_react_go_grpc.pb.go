// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: service_react_go.proto

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
	ReactGo_Hello_FullMethodName = "/pb.ReactGo/Hello"
)

// ReactGoClient is the client API for ReactGo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReactGoClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type reactGoClient struct {
	cc grpc.ClientConnInterface
}

func NewReactGoClient(cc grpc.ClientConnInterface) ReactGoClient {
	return &reactGoClient{cc}
}

func (c *reactGoClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, ReactGo_Hello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReactGoServer is the server API for ReactGo service.
// All implementations must embed UnimplementedReactGoServer
// for forward compatibility
type ReactGoServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	mustEmbedUnimplementedReactGoServer()
}

// UnimplementedReactGoServer must be embedded to have forward compatible implementations.
type UnimplementedReactGoServer struct {
}

func (UnimplementedReactGoServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedReactGoServer) mustEmbedUnimplementedReactGoServer() {}

// UnsafeReactGoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReactGoServer will
// result in compilation errors.
type UnsafeReactGoServer interface {
	mustEmbedUnimplementedReactGoServer()
}

func RegisterReactGoServer(s grpc.ServiceRegistrar, srv ReactGoServer) {
	s.RegisterService(&ReactGo_ServiceDesc, srv)
}

func _ReactGo_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReactGoServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReactGo_Hello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReactGoServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReactGo_ServiceDesc is the grpc.ServiceDesc for ReactGo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReactGo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ReactGo",
	HandlerType: (*ReactGoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _ReactGo_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_react_go.proto",
}
