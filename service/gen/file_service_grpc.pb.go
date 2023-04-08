// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: service/proto/file_service.proto

package gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	FileService_SaveFile_FullMethodName     = "/proto.FileService/SaveFile"
	FileService_GetFileList_FullMethodName  = "/proto.FileService/GetFileList"
	FileService_DownloadFile_FullMethodName = "/proto.FileService/DownloadFile"
)

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileServiceClient interface {
	SaveFile(ctx context.Context, in *SaveFileRequest, opts ...grpc.CallOption) (*SaveFileResponse, error)
	GetFileList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (FileService_GetFileListClient, error)
	DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (*DownloadFileResponse, error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) SaveFile(ctx context.Context, in *SaveFileRequest, opts ...grpc.CallOption) (*SaveFileResponse, error) {
	out := new(SaveFileResponse)
	err := c.cc.Invoke(ctx, FileService_SaveFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) GetFileList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (FileService_GetFileListClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileService_ServiceDesc.Streams[0], FileService_GetFileList_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &fileServiceGetFileListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileService_GetFileListClient interface {
	Recv() (*GetFileListResponse, error)
	grpc.ClientStream
}

type fileServiceGetFileListClient struct {
	grpc.ClientStream
}

func (x *fileServiceGetFileListClient) Recv() (*GetFileListResponse, error) {
	m := new(GetFileListResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileServiceClient) DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (*DownloadFileResponse, error) {
	out := new(DownloadFileResponse)
	err := c.cc.Invoke(ctx, FileService_DownloadFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServiceServer is the server API for FileService service.
// All implementations must embed UnimplementedFileServiceServer
// for forward compatibility
type FileServiceServer interface {
	SaveFile(context.Context, *SaveFileRequest) (*SaveFileResponse, error)
	GetFileList(*emptypb.Empty, FileService_GetFileListServer) error
	DownloadFile(context.Context, *DownloadFileRequest) (*DownloadFileResponse, error)
	mustEmbedUnimplementedFileServiceServer()
}

// UnimplementedFileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (UnimplementedFileServiceServer) SaveFile(context.Context, *SaveFileRequest) (*SaveFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveFile not implemented")
}
func (UnimplementedFileServiceServer) GetFileList(*emptypb.Empty, FileService_GetFileListServer) error {
	return status.Errorf(codes.Unimplemented, "method GetFileList not implemented")
}
func (UnimplementedFileServiceServer) DownloadFile(context.Context, *DownloadFileRequest) (*DownloadFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedFileServiceServer) mustEmbedUnimplementedFileServiceServer() {}

// UnsafeFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServiceServer will
// result in compilation errors.
type UnsafeFileServiceServer interface {
	mustEmbedUnimplementedFileServiceServer()
}

func RegisterFileServiceServer(s grpc.ServiceRegistrar, srv FileServiceServer) {
	s.RegisterService(&FileService_ServiceDesc, srv)
}

func _FileService_SaveFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).SaveFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_SaveFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).SaveFile(ctx, req.(*SaveFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_GetFileList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileServiceServer).GetFileList(m, &fileServiceGetFileListServer{stream})
}

type FileService_GetFileListServer interface {
	Send(*GetFileListResponse) error
	grpc.ServerStream
}

type fileServiceGetFileListServer struct {
	grpc.ServerStream
}

func (x *fileServiceGetFileListServer) Send(m *GetFileListResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _FileService_DownloadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).DownloadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_DownloadFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).DownloadFile(ctx, req.(*DownloadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FileService_ServiceDesc is the grpc.ServiceDesc for FileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveFile",
			Handler:    _FileService_SaveFile_Handler,
		},
		{
			MethodName: "DownloadFile",
			Handler:    _FileService_DownloadFile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetFileList",
			Handler:       _FileService_GetFileList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service/proto/file_service.proto",
}