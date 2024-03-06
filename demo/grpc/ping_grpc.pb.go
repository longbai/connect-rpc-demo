// Copyright 2021-2024 The Connect Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// The canonical location for this file is
// https://github.com/connectrpc/connect-go/blob/main/internal/proto/connect/ping/v1/ping.proto.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: ping.proto

// The connect.ping.v1 package contains an echo service designed to test the
// connect-go implementation.

package grpc

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
	PingService_Ping_FullMethodName    = "/connect.ping.v1.PingService/Ping"
	PingService_Fail_FullMethodName    = "/connect.ping.v1.PingService/Fail"
	PingService_Sum_FullMethodName     = "/connect.ping.v1.PingService/Sum"
	PingService_CountUp_FullMethodName = "/connect.ping.v1.PingService/CountUp"
	PingService_CumSum_FullMethodName  = "/connect.ping.v1.PingService/CumSum"
)

// PingServiceClient is the client API for PingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PingServiceClient interface {
	// Ping sends a ping to the server to determine if it's reachable.
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	// Fail always fails.
	Fail(ctx context.Context, in *FailRequest, opts ...grpc.CallOption) (*FailResponse, error)
	// Sum calculates the sum of the numbers sent on the stream.
	Sum(ctx context.Context, opts ...grpc.CallOption) (PingService_SumClient, error)
	// CountUp returns a stream of the numbers up to the given request.
	CountUp(ctx context.Context, in *CountUpRequest, opts ...grpc.CallOption) (PingService_CountUpClient, error)
	// CumSum determines the cumulative sum of all the numbers sent on the stream.
	CumSum(ctx context.Context, opts ...grpc.CallOption) (PingService_CumSumClient, error)
}

type pingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPingServiceClient(cc grpc.ClientConnInterface) PingServiceClient {
	return &pingServiceClient{cc}
}

func (c *pingServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, PingService_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pingServiceClient) Fail(ctx context.Context, in *FailRequest, opts ...grpc.CallOption) (*FailResponse, error) {
	out := new(FailResponse)
	err := c.cc.Invoke(ctx, PingService_Fail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pingServiceClient) Sum(ctx context.Context, opts ...grpc.CallOption) (PingService_SumClient, error) {
	stream, err := c.cc.NewStream(ctx, &PingService_ServiceDesc.Streams[0], PingService_Sum_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &pingServiceSumClient{stream}
	return x, nil
}

type PingService_SumClient interface {
	Send(*SumRequest) error
	CloseAndRecv() (*SumResponse, error)
	grpc.ClientStream
}

type pingServiceSumClient struct {
	grpc.ClientStream
}

func (x *pingServiceSumClient) Send(m *SumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pingServiceSumClient) CloseAndRecv() (*SumResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pingServiceClient) CountUp(ctx context.Context, in *CountUpRequest, opts ...grpc.CallOption) (PingService_CountUpClient, error) {
	stream, err := c.cc.NewStream(ctx, &PingService_ServiceDesc.Streams[1], PingService_CountUp_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &pingServiceCountUpClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PingService_CountUpClient interface {
	Recv() (*CountUpResponse, error)
	grpc.ClientStream
}

type pingServiceCountUpClient struct {
	grpc.ClientStream
}

func (x *pingServiceCountUpClient) Recv() (*CountUpResponse, error) {
	m := new(CountUpResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pingServiceClient) CumSum(ctx context.Context, opts ...grpc.CallOption) (PingService_CumSumClient, error) {
	stream, err := c.cc.NewStream(ctx, &PingService_ServiceDesc.Streams[2], PingService_CumSum_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &pingServiceCumSumClient{stream}
	return x, nil
}

type PingService_CumSumClient interface {
	Send(*CumSumRequest) error
	Recv() (*CumSumResponse, error)
	grpc.ClientStream
}

type pingServiceCumSumClient struct {
	grpc.ClientStream
}

func (x *pingServiceCumSumClient) Send(m *CumSumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pingServiceCumSumClient) Recv() (*CumSumResponse, error) {
	m := new(CumSumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PingServiceServer is the server API for PingService service.
// All implementations must embed UnimplementedPingServiceServer
// for forward compatibility
type PingServiceServer interface {
	// Ping sends a ping to the server to determine if it's reachable.
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	// Fail always fails.
	Fail(context.Context, *FailRequest) (*FailResponse, error)
	// Sum calculates the sum of the numbers sent on the stream.
	Sum(PingService_SumServer) error
	// CountUp returns a stream of the numbers up to the given request.
	CountUp(*CountUpRequest, PingService_CountUpServer) error
	// CumSum determines the cumulative sum of all the numbers sent on the stream.
	CumSum(PingService_CumSumServer) error
	mustEmbedUnimplementedPingServiceServer()
}

// UnimplementedPingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPingServiceServer struct {
}

func (UnimplementedPingServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedPingServiceServer) Fail(context.Context, *FailRequest) (*FailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fail not implemented")
}
func (UnimplementedPingServiceServer) Sum(PingService_SumServer) error {
	return status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedPingServiceServer) CountUp(*CountUpRequest, PingService_CountUpServer) error {
	return status.Errorf(codes.Unimplemented, "method CountUp not implemented")
}
func (UnimplementedPingServiceServer) CumSum(PingService_CumSumServer) error {
	return status.Errorf(codes.Unimplemented, "method CumSum not implemented")
}
func (UnimplementedPingServiceServer) mustEmbedUnimplementedPingServiceServer() {}

// UnsafePingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PingServiceServer will
// result in compilation errors.
type UnsafePingServiceServer interface {
	mustEmbedUnimplementedPingServiceServer()
}

func RegisterPingServiceServer(s grpc.ServiceRegistrar, srv PingServiceServer) {
	s.RegisterService(&PingService_ServiceDesc, srv)
}

func _PingService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PingService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PingService_Fail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingServiceServer).Fail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PingService_Fail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingServiceServer).Fail(ctx, req.(*FailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PingService_Sum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PingServiceServer).Sum(&pingServiceSumServer{stream})
}

type PingService_SumServer interface {
	SendAndClose(*SumResponse) error
	Recv() (*SumRequest, error)
	grpc.ServerStream
}

type pingServiceSumServer struct {
	grpc.ServerStream
}

func (x *pingServiceSumServer) SendAndClose(m *SumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pingServiceSumServer) Recv() (*SumRequest, error) {
	m := new(SumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PingService_CountUp_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CountUpRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PingServiceServer).CountUp(m, &pingServiceCountUpServer{stream})
}

type PingService_CountUpServer interface {
	Send(*CountUpResponse) error
	grpc.ServerStream
}

type pingServiceCountUpServer struct {
	grpc.ServerStream
}

func (x *pingServiceCountUpServer) Send(m *CountUpResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _PingService_CumSum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PingServiceServer).CumSum(&pingServiceCumSumServer{stream})
}

type PingService_CumSumServer interface {
	Send(*CumSumResponse) error
	Recv() (*CumSumRequest, error)
	grpc.ServerStream
}

type pingServiceCumSumServer struct {
	grpc.ServerStream
}

func (x *pingServiceCumSumServer) Send(m *CumSumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pingServiceCumSumServer) Recv() (*CumSumRequest, error) {
	m := new(CumSumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PingService_ServiceDesc is the grpc.ServiceDesc for PingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "connect.ping.v1.PingService",
	HandlerType: (*PingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _PingService_Ping_Handler,
		},
		{
			MethodName: "Fail",
			Handler:    _PingService_Fail_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Sum",
			Handler:       _PingService_Sum_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "CountUp",
			Handler:       _PingService_CountUp_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CumSum",
			Handler:       _PingService_CumSum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ping.proto",
}
