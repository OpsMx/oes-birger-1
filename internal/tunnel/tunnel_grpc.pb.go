//
// Copyright 2021-2023 OpsMx, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License")
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: internal/tunnel/tunnel.proto

package tunnel

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
	TunnelService_Hello_FullMethodName                     = "/tunnel.TunnelService/Hello"
	TunnelService_Ping_FullMethodName                      = "/tunnel.TunnelService/Ping"
	TunnelService_StartTunnel_FullMethodName               = "/tunnel.TunnelService/StartTunnel"
	TunnelService_WaitForRequest_FullMethodName            = "/tunnel.TunnelService/WaitForRequest"
	TunnelService_DataFlowAgentToController_FullMethodName = "/tunnel.TunnelService/DataFlowAgentToController"
)

// TunnelServiceClient is the client API for TunnelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TunnelServiceClient interface {
	// Initial handshake, sent from agent to controller on the control connection, before any other messages.
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	// Keep alive, sent from the agent to the controller, on control and data connections.
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	// The agent will call StartRequest() when it wants the controller to perform
	// an HTTP fetch.
	StartTunnel(ctx context.Context, in *TunnelRequest, opts ...grpc.CallOption) (*TunnelHeaders, error)
	// The agent will perform a long-running call to WaitForRequest() and handle
	// any HTTP request found.
	WaitForRequest(ctx context.Context, in *WaitForRequestArgs, opts ...grpc.CallOption) (TunnelService_WaitForRequestClient, error)
	// DataFlowAgentToController is the conduit for an agent to send and manage the
	// flow of data for a specific HTTP response.
	DataFlowAgentToController(ctx context.Context, opts ...grpc.CallOption) (TunnelService_DataFlowAgentToControllerClient, error)
}

type tunnelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTunnelServiceClient(cc grpc.ClientConnInterface) TunnelServiceClient {
	return &tunnelServiceClient{cc}
}

func (c *tunnelServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, TunnelService_Hello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tunnelServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, TunnelService_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tunnelServiceClient) StartTunnel(ctx context.Context, in *TunnelRequest, opts ...grpc.CallOption) (*TunnelHeaders, error) {
	out := new(TunnelHeaders)
	err := c.cc.Invoke(ctx, TunnelService_StartTunnel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tunnelServiceClient) WaitForRequest(ctx context.Context, in *WaitForRequestArgs, opts ...grpc.CallOption) (TunnelService_WaitForRequestClient, error) {
	stream, err := c.cc.NewStream(ctx, &TunnelService_ServiceDesc.Streams[0], TunnelService_WaitForRequest_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &tunnelServiceWaitForRequestClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TunnelService_WaitForRequestClient interface {
	Recv() (*TunnelRequest, error)
	grpc.ClientStream
}

type tunnelServiceWaitForRequestClient struct {
	grpc.ClientStream
}

func (x *tunnelServiceWaitForRequestClient) Recv() (*TunnelRequest, error) {
	m := new(TunnelRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tunnelServiceClient) DataFlowAgentToController(ctx context.Context, opts ...grpc.CallOption) (TunnelService_DataFlowAgentToControllerClient, error) {
	stream, err := c.cc.NewStream(ctx, &TunnelService_ServiceDesc.Streams[1], TunnelService_DataFlowAgentToController_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &tunnelServiceDataFlowAgentToControllerClient{stream}
	return x, nil
}

type TunnelService_DataFlowAgentToControllerClient interface {
	Send(*StreamFlow) error
	CloseAndRecv() (*StreamFlowResponse, error)
	grpc.ClientStream
}

type tunnelServiceDataFlowAgentToControllerClient struct {
	grpc.ClientStream
}

func (x *tunnelServiceDataFlowAgentToControllerClient) Send(m *StreamFlow) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tunnelServiceDataFlowAgentToControllerClient) CloseAndRecv() (*StreamFlowResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamFlowResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TunnelServiceServer is the server API for TunnelService service.
// All implementations must embed UnimplementedTunnelServiceServer
// for forward compatibility
type TunnelServiceServer interface {
	// Initial handshake, sent from agent to controller on the control connection, before any other messages.
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	// Keep alive, sent from the agent to the controller, on control and data connections.
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	// The agent will call StartRequest() when it wants the controller to perform
	// an HTTP fetch.
	StartTunnel(context.Context, *TunnelRequest) (*TunnelHeaders, error)
	// The agent will perform a long-running call to WaitForRequest() and handle
	// any HTTP request found.
	WaitForRequest(*WaitForRequestArgs, TunnelService_WaitForRequestServer) error
	// DataFlowAgentToController is the conduit for an agent to send and manage the
	// flow of data for a specific HTTP response.
	DataFlowAgentToController(TunnelService_DataFlowAgentToControllerServer) error
	mustEmbedUnimplementedTunnelServiceServer()
}

// UnimplementedTunnelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTunnelServiceServer struct {
}

func (UnimplementedTunnelServiceServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedTunnelServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedTunnelServiceServer) StartTunnel(context.Context, *TunnelRequest) (*TunnelHeaders, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartTunnel not implemented")
}
func (UnimplementedTunnelServiceServer) WaitForRequest(*WaitForRequestArgs, TunnelService_WaitForRequestServer) error {
	return status.Errorf(codes.Unimplemented, "method WaitForRequest not implemented")
}
func (UnimplementedTunnelServiceServer) DataFlowAgentToController(TunnelService_DataFlowAgentToControllerServer) error {
	return status.Errorf(codes.Unimplemented, "method DataFlowAgentToController not implemented")
}
func (UnimplementedTunnelServiceServer) mustEmbedUnimplementedTunnelServiceServer() {}

// UnsafeTunnelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TunnelServiceServer will
// result in compilation errors.
type UnsafeTunnelServiceServer interface {
	mustEmbedUnimplementedTunnelServiceServer()
}

func RegisterTunnelServiceServer(s grpc.ServiceRegistrar, srv TunnelServiceServer) {
	s.RegisterService(&TunnelService_ServiceDesc, srv)
}

func _TunnelService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TunnelServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TunnelService_Hello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TunnelServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TunnelService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TunnelServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TunnelService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TunnelServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TunnelService_StartTunnel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TunnelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TunnelServiceServer).StartTunnel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TunnelService_StartTunnel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TunnelServiceServer).StartTunnel(ctx, req.(*TunnelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TunnelService_WaitForRequest_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WaitForRequestArgs)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TunnelServiceServer).WaitForRequest(m, &tunnelServiceWaitForRequestServer{stream})
}

type TunnelService_WaitForRequestServer interface {
	Send(*TunnelRequest) error
	grpc.ServerStream
}

type tunnelServiceWaitForRequestServer struct {
	grpc.ServerStream
}

func (x *tunnelServiceWaitForRequestServer) Send(m *TunnelRequest) error {
	return x.ServerStream.SendMsg(m)
}

func _TunnelService_DataFlowAgentToController_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TunnelServiceServer).DataFlowAgentToController(&tunnelServiceDataFlowAgentToControllerServer{stream})
}

type TunnelService_DataFlowAgentToControllerServer interface {
	SendAndClose(*StreamFlowResponse) error
	Recv() (*StreamFlow, error)
	grpc.ServerStream
}

type tunnelServiceDataFlowAgentToControllerServer struct {
	grpc.ServerStream
}

func (x *tunnelServiceDataFlowAgentToControllerServer) SendAndClose(m *StreamFlowResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tunnelServiceDataFlowAgentToControllerServer) Recv() (*StreamFlow, error) {
	m := new(StreamFlow)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TunnelService_ServiceDesc is the grpc.ServiceDesc for TunnelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TunnelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tunnel.TunnelService",
	HandlerType: (*TunnelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _TunnelService_Hello_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _TunnelService_Ping_Handler,
		},
		{
			MethodName: "StartTunnel",
			Handler:    _TunnelService_StartTunnel_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WaitForRequest",
			Handler:       _TunnelService_WaitForRequest_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DataFlowAgentToController",
			Handler:       _TunnelService_DataFlowAgentToController_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "internal/tunnel/tunnel.proto",
}
