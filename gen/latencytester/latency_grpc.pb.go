// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/latency.proto

package latencytester

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	LatencyService_MeasureLatency_FullMethodName = "/latency.LatencyService/MeasureLatency"
)

// LatencyServiceClient is the client API for LatencyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LatencyServiceClient interface {
	MeasureLatency(ctx context.Context, in *LatencyRequest, opts ...grpc.CallOption) (*LatencyResponse, error)
}

type latencyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLatencyServiceClient(cc grpc.ClientConnInterface) LatencyServiceClient {
	return &latencyServiceClient{cc}
}

func (c *latencyServiceClient) MeasureLatency(ctx context.Context, in *LatencyRequest, opts ...grpc.CallOption) (*LatencyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LatencyResponse)
	err := c.cc.Invoke(ctx, LatencyService_MeasureLatency_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LatencyServiceServer is the server API for LatencyService service.
// All implementations must embed UnimplementedLatencyServiceServer
// for forward compatibility.
type LatencyServiceServer interface {
	MeasureLatency(context.Context, *LatencyRequest) (*LatencyResponse, error)
	mustEmbedUnimplementedLatencyServiceServer()
}

// UnimplementedLatencyServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLatencyServiceServer struct{}

func (UnimplementedLatencyServiceServer) MeasureLatency(context.Context, *LatencyRequest) (*LatencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MeasureLatency not implemented")
}
func (UnimplementedLatencyServiceServer) mustEmbedUnimplementedLatencyServiceServer() {}
func (UnimplementedLatencyServiceServer) testEmbeddedByValue()                        {}

// UnsafeLatencyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LatencyServiceServer will
// result in compilation errors.
type UnsafeLatencyServiceServer interface {
	mustEmbedUnimplementedLatencyServiceServer()
}

func RegisterLatencyServiceServer(s grpc.ServiceRegistrar, srv LatencyServiceServer) {
	// If the following call pancis, it indicates UnimplementedLatencyServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LatencyService_ServiceDesc, srv)
}

func _LatencyService_MeasureLatency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LatencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LatencyServiceServer).MeasureLatency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LatencyService_MeasureLatency_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LatencyServiceServer).MeasureLatency(ctx, req.(*LatencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LatencyService_ServiceDesc is the grpc.ServiceDesc for LatencyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LatencyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "latency.LatencyService",
	HandlerType: (*LatencyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MeasureLatency",
			Handler:    _LatencyService_MeasureLatency_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/latency.proto",
}
