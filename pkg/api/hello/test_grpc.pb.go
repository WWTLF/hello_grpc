// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.3
// source: api/hello/test.proto

package hello

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

// PingClient is the client API for Ping service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PingClient interface {
	SayHello(ctx context.Context, in *Test, opts ...grpc.CallOption) (Ping_SayHelloClient, error)
}

type pingClient struct {
	cc grpc.ClientConnInterface
}

func NewPingClient(cc grpc.ClientConnInterface) PingClient {
	return &pingClient{cc}
}

func (c *pingClient) SayHello(ctx context.Context, in *Test, opts ...grpc.CallOption) (Ping_SayHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &Ping_ServiceDesc.Streams[0], "/wwtlf.v1.Ping/SayHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &pingSayHelloClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Ping_SayHelloClient interface {
	Recv() (*Test, error)
	grpc.ClientStream
}

type pingSayHelloClient struct {
	grpc.ClientStream
}

func (x *pingSayHelloClient) Recv() (*Test, error) {
	m := new(Test)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PingServer is the server API for Ping service.
// All implementations must embed UnimplementedPingServer
// for forward compatibility
type PingServer interface {
	SayHello(*Test, Ping_SayHelloServer) error
	mustEmbedUnimplementedPingServer()
}

// UnimplementedPingServer must be embedded to have forward compatible implementations.
type UnimplementedPingServer struct {
}

func (UnimplementedPingServer) SayHello(*Test, Ping_SayHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedPingServer) mustEmbedUnimplementedPingServer() {}

// UnsafePingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PingServer will
// result in compilation errors.
type UnsafePingServer interface {
	mustEmbedUnimplementedPingServer()
}

func RegisterPingServer(s grpc.ServiceRegistrar, srv PingServer) {
	s.RegisterService(&Ping_ServiceDesc, srv)
}

func _Ping_SayHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Test)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PingServer).SayHello(m, &pingSayHelloServer{stream})
}

type Ping_SayHelloServer interface {
	Send(*Test) error
	grpc.ServerStream
}

type pingSayHelloServer struct {
	grpc.ServerStream
}

func (x *pingSayHelloServer) Send(m *Test) error {
	return x.ServerStream.SendMsg(m)
}

// Ping_ServiceDesc is the grpc.ServiceDesc for Ping service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ping_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wwtlf.v1.Ping",
	HandlerType: (*PingServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHello",
			Handler:       _Ping_SayHello_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/hello/test.proto",
}