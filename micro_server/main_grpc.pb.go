// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: micro_server/main.proto

package micro_server_package

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	services "vvfan/grpc_demo/micro_server/services"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AhaResourcesClient is the client API for AhaResources service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AhaResourcesClient interface {
	Add(ctx context.Context, in *services.TestParams, opts ...grpc.CallOption) (*services.TestResponse, error)
}

type ahaResourcesClient struct {
	cc grpc.ClientConnInterface
}

func NewAhaResourcesClient(cc grpc.ClientConnInterface) AhaResourcesClient {
	return &ahaResourcesClient{cc}
}

func (c *ahaResourcesClient) Add(ctx context.Context, in *services.TestParams, opts ...grpc.CallOption) (*services.TestResponse, error) {
	out := new(services.TestResponse)
	err := c.cc.Invoke(ctx, "/micro_server.aha_resources/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AhaResourcesServer is the server API for AhaResources service.
// All implementations must embed UnimplementedAhaResourcesServer
// for forward compatibility
type AhaResourcesServer interface {
	Add(context.Context, *services.TestParams) (*services.TestResponse, error)
	mustEmbedUnimplementedAhaResourcesServer()
}

// UnimplementedAhaResourcesServer must be embedded to have forward compatible implementations.
type UnimplementedAhaResourcesServer struct {
}

func (UnimplementedAhaResourcesServer) Add(context.Context, *services.TestParams) (*services.TestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedAhaResourcesServer) mustEmbedUnimplementedAhaResourcesServer() {}

// UnsafeAhaResourcesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AhaResourcesServer will
// result in compilation errors.
type UnsafeAhaResourcesServer interface {
	mustEmbedUnimplementedAhaResourcesServer()
}

func RegisterAhaResourcesServer(s grpc.ServiceRegistrar, srv AhaResourcesServer) {
	s.RegisterService(&AhaResources_ServiceDesc, srv)
}

func _AhaResources_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(services.TestParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AhaResourcesServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/micro_server.aha_resources/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AhaResourcesServer).Add(ctx, req.(*services.TestParams))
	}
	return interceptor(ctx, in, info, handler)
}

// AhaResources_ServiceDesc is the grpc.ServiceDesc for AhaResources service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AhaResources_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "micro_server.aha_resources",
	HandlerType: (*AhaResourcesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _AhaResources_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "micro_server/main.proto",
}
