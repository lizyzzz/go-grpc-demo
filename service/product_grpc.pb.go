// protoc --go_out=./ --go-grpc_out=./ ./product.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: product.proto

// 生成消息的包名

package service

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
	ProdService_GetProductStock_FullMethodName                = "/service.ProdService/GetProductStock"
	ProdService_UpdateProductStockClientStream_FullMethodName = "/service.ProdService/UpdateProductStockClientStream"
	ProdService_GetProductStockServerStream_FullMethodName    = "/service.ProdService/GetProductStockServerStream"
	ProdService_SayHelloStream_FullMethodName                 = "/service.ProdService/SayHelloStream"
)

// ProdServiceClient is the client API for ProdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProdServiceClient interface {
	// rpc 服务的函数名 (传入参数) 返回 (返回参数)
	GetProductStock(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
	// 客户端流
	// 客户端发送多个请求, 服务端只回复一个响应
	UpdateProductStockClientStream(ctx context.Context, opts ...grpc.CallOption) (ProdService_UpdateProductStockClientStreamClient, error)
	// 服务端流
	// 客户端发送一个请求, 服务端回复多个响应
	GetProductStockServerStream(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (ProdService_GetProductStockServerStreamClient, error)
	// 双向流
	// 客户端和服务端都能发多个消息, 通信模式
	SayHelloStream(ctx context.Context, opts ...grpc.CallOption) (ProdService_SayHelloStreamClient, error)
}

type prodServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProdServiceClient(cc grpc.ClientConnInterface) ProdServiceClient {
	return &prodServiceClient{cc}
}

func (c *prodServiceClient) GetProductStock(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, ProdService_GetProductStock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prodServiceClient) UpdateProductStockClientStream(ctx context.Context, opts ...grpc.CallOption) (ProdService_UpdateProductStockClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProdService_ServiceDesc.Streams[0], ProdService_UpdateProductStockClientStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &prodServiceUpdateProductStockClientStreamClient{stream}
	return x, nil
}

type ProdService_UpdateProductStockClientStreamClient interface {
	Send(*ProductRequest) error
	CloseAndRecv() (*ProductResponse, error)
	grpc.ClientStream
}

type prodServiceUpdateProductStockClientStreamClient struct {
	grpc.ClientStream
}

func (x *prodServiceUpdateProductStockClientStreamClient) Send(m *ProductRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *prodServiceUpdateProductStockClientStreamClient) CloseAndRecv() (*ProductResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ProductResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *prodServiceClient) GetProductStockServerStream(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (ProdService_GetProductStockServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProdService_ServiceDesc.Streams[1], ProdService_GetProductStockServerStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &prodServiceGetProductStockServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProdService_GetProductStockServerStreamClient interface {
	Recv() (*ProductResponse, error)
	grpc.ClientStream
}

type prodServiceGetProductStockServerStreamClient struct {
	grpc.ClientStream
}

func (x *prodServiceGetProductStockServerStreamClient) Recv() (*ProductResponse, error) {
	m := new(ProductResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *prodServiceClient) SayHelloStream(ctx context.Context, opts ...grpc.CallOption) (ProdService_SayHelloStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProdService_ServiceDesc.Streams[2], ProdService_SayHelloStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &prodServiceSayHelloStreamClient{stream}
	return x, nil
}

type ProdService_SayHelloStreamClient interface {
	Send(*ClientMsg) error
	Recv() (*ServerMsg, error)
	grpc.ClientStream
}

type prodServiceSayHelloStreamClient struct {
	grpc.ClientStream
}

func (x *prodServiceSayHelloStreamClient) Send(m *ClientMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *prodServiceSayHelloStreamClient) Recv() (*ServerMsg, error) {
	m := new(ServerMsg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProdServiceServer is the server API for ProdService service.
// All implementations must embed UnimplementedProdServiceServer
// for forward compatibility
type ProdServiceServer interface {
	// rpc 服务的函数名 (传入参数) 返回 (返回参数)
	GetProductStock(context.Context, *ProductRequest) (*ProductResponse, error)
	// 客户端流
	// 客户端发送多个请求, 服务端只回复一个响应
	UpdateProductStockClientStream(ProdService_UpdateProductStockClientStreamServer) error
	// 服务端流
	// 客户端发送一个请求, 服务端回复多个响应
	GetProductStockServerStream(*ProductRequest, ProdService_GetProductStockServerStreamServer) error
	// 双向流
	// 客户端和服务端都能发多个消息, 通信模式
	SayHelloStream(ProdService_SayHelloStreamServer) error
	mustEmbedUnimplementedProdServiceServer()
}

// UnimplementedProdServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProdServiceServer struct {
}

func (UnimplementedProdServiceServer) GetProductStock(context.Context, *ProductRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductStock not implemented")
}
func (UnimplementedProdServiceServer) UpdateProductStockClientStream(ProdService_UpdateProductStockClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateProductStockClientStream not implemented")
}
func (UnimplementedProdServiceServer) GetProductStockServerStream(*ProductRequest, ProdService_GetProductStockServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetProductStockServerStream not implemented")
}
func (UnimplementedProdServiceServer) SayHelloStream(ProdService_SayHelloStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloStream not implemented")
}
func (UnimplementedProdServiceServer) mustEmbedUnimplementedProdServiceServer() {}

// UnsafeProdServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProdServiceServer will
// result in compilation errors.
type UnsafeProdServiceServer interface {
	mustEmbedUnimplementedProdServiceServer()
}

func RegisterProdServiceServer(s grpc.ServiceRegistrar, srv ProdServiceServer) {
	s.RegisterService(&ProdService_ServiceDesc, srv)
}

func _ProdService_GetProductStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProdServiceServer).GetProductStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProdService_GetProductStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProdServiceServer).GetProductStock(ctx, req.(*ProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProdService_UpdateProductStockClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProdServiceServer).UpdateProductStockClientStream(&prodServiceUpdateProductStockClientStreamServer{stream})
}

type ProdService_UpdateProductStockClientStreamServer interface {
	SendAndClose(*ProductResponse) error
	Recv() (*ProductRequest, error)
	grpc.ServerStream
}

type prodServiceUpdateProductStockClientStreamServer struct {
	grpc.ServerStream
}

func (x *prodServiceUpdateProductStockClientStreamServer) SendAndClose(m *ProductResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *prodServiceUpdateProductStockClientStreamServer) Recv() (*ProductRequest, error) {
	m := new(ProductRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ProdService_GetProductStockServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProductRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProdServiceServer).GetProductStockServerStream(m, &prodServiceGetProductStockServerStreamServer{stream})
}

type ProdService_GetProductStockServerStreamServer interface {
	Send(*ProductResponse) error
	grpc.ServerStream
}

type prodServiceGetProductStockServerStreamServer struct {
	grpc.ServerStream
}

func (x *prodServiceGetProductStockServerStreamServer) Send(m *ProductResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ProdService_SayHelloStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProdServiceServer).SayHelloStream(&prodServiceSayHelloStreamServer{stream})
}

type ProdService_SayHelloStreamServer interface {
	Send(*ServerMsg) error
	Recv() (*ClientMsg, error)
	grpc.ServerStream
}

type prodServiceSayHelloStreamServer struct {
	grpc.ServerStream
}

func (x *prodServiceSayHelloStreamServer) Send(m *ServerMsg) error {
	return x.ServerStream.SendMsg(m)
}

func (x *prodServiceSayHelloStreamServer) Recv() (*ClientMsg, error) {
	m := new(ClientMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProdService_ServiceDesc is the grpc.ServiceDesc for ProdService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProdService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.ProdService",
	HandlerType: (*ProdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProductStock",
			Handler:    _ProdService_GetProductStock_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UpdateProductStockClientStream",
			Handler:       _ProdService_UpdateProductStockClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetProductStockServerStream",
			Handler:       _ProdService_GetProductStockServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SayHelloStream",
			Handler:       _ProdService_SayHelloStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "product.proto",
}
