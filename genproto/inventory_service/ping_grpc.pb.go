// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: ping.proto

package inventory_service

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

// InventoryServiceClient is the client API for InventoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InventoryServiceClient interface {
	Ping(ctx context.Context, in *PingPong, opts ...grpc.CallOption) (*PingPong, error)
}

type inventoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryServiceClient(cc grpc.ClientConnInterface) InventoryServiceClient {
	return &inventoryServiceClient{cc}
}

func (c *inventoryServiceClient) Ping(ctx context.Context, in *PingPong, opts ...grpc.CallOption) (*PingPong, error) {
	out := new(PingPong)
	err := c.cc.Invoke(ctx, "/InventoryService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryServiceServer is the server API for InventoryService service.
// All implementations must embed UnimplementedInventoryServiceServer
// for forward compatibility
type InventoryServiceServer interface {
	Ping(context.Context, *PingPong) (*PingPong, error)
	mustEmbedUnimplementedInventoryServiceServer()
}

// UnimplementedInventoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInventoryServiceServer struct {
}

func (UnimplementedInventoryServiceServer) Ping(context.Context, *PingPong) (*PingPong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedInventoryServiceServer) mustEmbedUnimplementedInventoryServiceServer() {}

// UnsafeInventoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InventoryServiceServer will
// result in compilation errors.
type UnsafeInventoryServiceServer interface {
	mustEmbedUnimplementedInventoryServiceServer()
}

func RegisterInventoryServiceServer(s grpc.ServiceRegistrar, srv InventoryServiceServer) {
	s.RegisterService(&InventoryService_ServiceDesc, srv)
}

func _InventoryService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingPong)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InventoryService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).Ping(ctx, req.(*PingPong))
	}
	return interceptor(ctx, in, info, handler)
}

// InventoryService_ServiceDesc is the grpc.ServiceDesc for InventoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InventoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "InventoryService",
	HandlerType: (*InventoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _InventoryService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ping.proto",
}
