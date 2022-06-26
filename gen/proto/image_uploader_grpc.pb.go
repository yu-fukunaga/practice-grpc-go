// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: proto/image_uploader.proto

package proto

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

// ImageUploadServiceClient is the client API for ImageUploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageUploadServiceClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (ImageUploadService_UploadClient, error)
}

type imageUploadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImageUploadServiceClient(cc grpc.ClientConnInterface) ImageUploadServiceClient {
	return &imageUploadServiceClient{cc}
}

func (c *imageUploadServiceClient) Upload(ctx context.Context, opts ...grpc.CallOption) (ImageUploadService_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &ImageUploadService_ServiceDesc.Streams[0], "/image.uploader.ImageUploadService/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &imageUploadServiceUploadClient{stream}
	return x, nil
}

type ImageUploadService_UploadClient interface {
	Send(*ImageUploadRequest) error
	CloseAndRecv() (*ImageUploadResponse, error)
	grpc.ClientStream
}

type imageUploadServiceUploadClient struct {
	grpc.ClientStream
}

func (x *imageUploadServiceUploadClient) Send(m *ImageUploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *imageUploadServiceUploadClient) CloseAndRecv() (*ImageUploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ImageUploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ImageUploadServiceServer is the server API for ImageUploadService service.
// All implementations must embed UnimplementedImageUploadServiceServer
// for forward compatibility
type ImageUploadServiceServer interface {
	Upload(ImageUploadService_UploadServer) error
	mustEmbedUnimplementedImageUploadServiceServer()
}

// UnimplementedImageUploadServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImageUploadServiceServer struct {
}

func (UnimplementedImageUploadServiceServer) Upload(ImageUploadService_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedImageUploadServiceServer) mustEmbedUnimplementedImageUploadServiceServer() {}

// UnsafeImageUploadServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageUploadServiceServer will
// result in compilation errors.
type UnsafeImageUploadServiceServer interface {
	mustEmbedUnimplementedImageUploadServiceServer()
}

func RegisterImageUploadServiceServer(s grpc.ServiceRegistrar, srv ImageUploadServiceServer) {
	s.RegisterService(&ImageUploadService_ServiceDesc, srv)
}

func _ImageUploadService_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ImageUploadServiceServer).Upload(&imageUploadServiceUploadServer{stream})
}

type ImageUploadService_UploadServer interface {
	SendAndClose(*ImageUploadResponse) error
	Recv() (*ImageUploadRequest, error)
	grpc.ServerStream
}

type imageUploadServiceUploadServer struct {
	grpc.ServerStream
}

func (x *imageUploadServiceUploadServer) SendAndClose(m *ImageUploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *imageUploadServiceUploadServer) Recv() (*ImageUploadRequest, error) {
	m := new(ImageUploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ImageUploadService_ServiceDesc is the grpc.ServiceDesc for ImageUploadService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageUploadService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "image.uploader.ImageUploadService",
	HandlerType: (*ImageUploadServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _ImageUploadService_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/image_uploader.proto",
}
