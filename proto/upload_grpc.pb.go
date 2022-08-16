// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: proto/upload.proto

package uploadpb

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

// UploadServiceClient is the client API for UploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UploadServiceClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (UploadService_UploadClient, error)
	Send(ctx context.Context, opts ...grpc.CallOption) (UploadService_SendClient, error)
}

type uploadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUploadServiceClient(cc grpc.ClientConnInterface) UploadServiceClient {
	return &uploadServiceClient{cc}
}

func (c *uploadServiceClient) Upload(ctx context.Context, opts ...grpc.CallOption) (UploadService_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &UploadService_ServiceDesc.Streams[0], "/proto.UploadService/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &uploadServiceUploadClient{stream}
	return x, nil
}

type UploadService_UploadClient interface {
	Send(*UploadRequest) error
	CloseAndRecv() (*UploadResponse, error)
	grpc.ClientStream
}

type uploadServiceUploadClient struct {
	grpc.ClientStream
}

func (x *uploadServiceUploadClient) Send(m *UploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uploadServiceUploadClient) CloseAndRecv() (*UploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *uploadServiceClient) Send(ctx context.Context, opts ...grpc.CallOption) (UploadService_SendClient, error) {
	stream, err := c.cc.NewStream(ctx, &UploadService_ServiceDesc.Streams[1], "/proto.UploadService/Send", opts...)
	if err != nil {
		return nil, err
	}
	x := &uploadServiceSendClient{stream}
	return x, nil
}

type UploadService_SendClient interface {
	Send(*UploadRequest) error
	CloseAndRecv() (*UploadResponse, error)
	grpc.ClientStream
}

type uploadServiceSendClient struct {
	grpc.ClientStream
}

func (x *uploadServiceSendClient) Send(m *UploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uploadServiceSendClient) CloseAndRecv() (*UploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UploadServiceServer is the server API for UploadService service.
// All implementations must embed UnimplementedUploadServiceServer
// for forward compatibility
type UploadServiceServer interface {
	Upload(UploadService_UploadServer) error
	Send(UploadService_SendServer) error
	mustEmbedUnimplementedUploadServiceServer()
}

// UnimplementedUploadServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUploadServiceServer struct {
}

func (UnimplementedUploadServiceServer) Upload(UploadService_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedUploadServiceServer) Send(UploadService_SendServer) error {
	return status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedUploadServiceServer) mustEmbedUnimplementedUploadServiceServer() {}

// UnsafeUploadServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UploadServiceServer will
// result in compilation errors.
type UnsafeUploadServiceServer interface {
	mustEmbedUnimplementedUploadServiceServer()
}

func RegisterUploadServiceServer(s grpc.ServiceRegistrar, srv UploadServiceServer) {
	s.RegisterService(&UploadService_ServiceDesc, srv)
}

func _UploadService_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UploadServiceServer).Upload(&uploadServiceUploadServer{stream})
}

type UploadService_UploadServer interface {
	SendAndClose(*UploadResponse) error
	Recv() (*UploadRequest, error)
	grpc.ServerStream
}

type uploadServiceUploadServer struct {
	grpc.ServerStream
}

func (x *uploadServiceUploadServer) SendAndClose(m *UploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uploadServiceUploadServer) Recv() (*UploadRequest, error) {
	m := new(UploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _UploadService_Send_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UploadServiceServer).Send(&uploadServiceSendServer{stream})
}

type UploadService_SendServer interface {
	SendAndClose(*UploadResponse) error
	Recv() (*UploadRequest, error)
	grpc.ServerStream
}

type uploadServiceSendServer struct {
	grpc.ServerStream
}

func (x *uploadServiceSendServer) SendAndClose(m *UploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uploadServiceSendServer) Recv() (*UploadRequest, error) {
	m := new(UploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UploadService_ServiceDesc is the grpc.ServiceDesc for UploadService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UploadService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UploadService",
	HandlerType: (*UploadServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _UploadService_Upload_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Send",
			Handler:       _UploadService_Send_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/upload.proto",
}
