// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.3
// source: proto/codiocmd/codiocmd_service.proto

package codiocmd

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
	CodioCmd_FileUpload_FullMethodName        = "/com.codio.chipmunk.proto.codiocmd.CodioCmd/FileUpload"
	CodioCmd_FileDownload_FullMethodName      = "/com.codio.chipmunk.proto.codiocmd.CodioCmd/FileDownload"
	CodioCmd_Exec_FullMethodName              = "/com.codio.chipmunk.proto.codiocmd.CodioCmd/Exec"
	CodioCmd_ExecAsync_FullMethodName         = "/com.codio.chipmunk.proto.codiocmd.CodioCmd/ExecAsync"
	CodioCmd_Ping_FullMethodName              = "/com.codio.chipmunk.proto.codiocmd.CodioCmd/Ping"
	CodioCmd_GetVMHostName_FullMethodName     = "/com.codio.chipmunk.proto.codiocmd.CodioCmd/GetVMHostName"
	CodioCmd_GetRemoteFileList_FullMethodName = "/com.codio.chipmunk.proto.codiocmd.CodioCmd/GetRemoteFileList"
	CodioCmd_SyncPath_FullMethodName          = "/com.codio.chipmunk.proto.codiocmd.CodioCmd/SyncPath"
	CodioCmd_GetReusableToken_FullMethodName  = "/com.codio.chipmunk.proto.codiocmd.CodioCmd/GetReusableToken"
)

// CodioCmdClient is the client API for CodioCmd service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CodioCmdClient interface {
	FileUpload(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[FileUploadRequest, FileUploadResponse], error)
	FileDownload(ctx context.Context, in *FileDownloadRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[FileDownloadResponse], error)
	Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecResponse, error)
	ExecAsync(ctx context.Context, in *ExecAsyncRequest, opts ...grpc.CallOption) (*ExecAsyncResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	GetVMHostName(ctx context.Context, in *GetVMHostNameRequest, opts ...grpc.CallOption) (*GetVMHostNameResponse, error)
	GetRemoteFileList(ctx context.Context, in *GetRemoteFileListRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetRemoteFileListResponse], error)
	SyncPath(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[FileWatcherEvent, FileWatcherEvent], error)
	GetReusableToken(ctx context.Context, in *GetReusableTokenRequest, opts ...grpc.CallOption) (*GetReusableTokenResponse, error)
}

type codioCmdClient struct {
	cc grpc.ClientConnInterface
}

func NewCodioCmdClient(cc grpc.ClientConnInterface) CodioCmdClient {
	return &codioCmdClient{cc}
}

func (c *codioCmdClient) FileUpload(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[FileUploadRequest, FileUploadResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CodioCmd_ServiceDesc.Streams[0], CodioCmd_FileUpload_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[FileUploadRequest, FileUploadResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CodioCmd_FileUploadClient = grpc.ClientStreamingClient[FileUploadRequest, FileUploadResponse]

func (c *codioCmdClient) FileDownload(ctx context.Context, in *FileDownloadRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[FileDownloadResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CodioCmd_ServiceDesc.Streams[1], CodioCmd_FileDownload_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[FileDownloadRequest, FileDownloadResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CodioCmd_FileDownloadClient = grpc.ServerStreamingClient[FileDownloadResponse]

func (c *codioCmdClient) Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExecResponse)
	err := c.cc.Invoke(ctx, CodioCmd_Exec_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codioCmdClient) ExecAsync(ctx context.Context, in *ExecAsyncRequest, opts ...grpc.CallOption) (*ExecAsyncResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExecAsyncResponse)
	err := c.cc.Invoke(ctx, CodioCmd_ExecAsync_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codioCmdClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, CodioCmd_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codioCmdClient) GetVMHostName(ctx context.Context, in *GetVMHostNameRequest, opts ...grpc.CallOption) (*GetVMHostNameResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVMHostNameResponse)
	err := c.cc.Invoke(ctx, CodioCmd_GetVMHostName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codioCmdClient) GetRemoteFileList(ctx context.Context, in *GetRemoteFileListRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetRemoteFileListResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CodioCmd_ServiceDesc.Streams[2], CodioCmd_GetRemoteFileList_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetRemoteFileListRequest, GetRemoteFileListResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CodioCmd_GetRemoteFileListClient = grpc.ServerStreamingClient[GetRemoteFileListResponse]

func (c *codioCmdClient) SyncPath(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[FileWatcherEvent, FileWatcherEvent], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CodioCmd_ServiceDesc.Streams[3], CodioCmd_SyncPath_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[FileWatcherEvent, FileWatcherEvent]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CodioCmd_SyncPathClient = grpc.BidiStreamingClient[FileWatcherEvent, FileWatcherEvent]

func (c *codioCmdClient) GetReusableToken(ctx context.Context, in *GetReusableTokenRequest, opts ...grpc.CallOption) (*GetReusableTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetReusableTokenResponse)
	err := c.cc.Invoke(ctx, CodioCmd_GetReusableToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CodioCmdServer is the server API for CodioCmd service.
// All implementations must embed UnimplementedCodioCmdServer
// for forward compatibility.
type CodioCmdServer interface {
	FileUpload(grpc.ClientStreamingServer[FileUploadRequest, FileUploadResponse]) error
	FileDownload(*FileDownloadRequest, grpc.ServerStreamingServer[FileDownloadResponse]) error
	Exec(context.Context, *ExecRequest) (*ExecResponse, error)
	ExecAsync(context.Context, *ExecAsyncRequest) (*ExecAsyncResponse, error)
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	GetVMHostName(context.Context, *GetVMHostNameRequest) (*GetVMHostNameResponse, error)
	GetRemoteFileList(*GetRemoteFileListRequest, grpc.ServerStreamingServer[GetRemoteFileListResponse]) error
	SyncPath(grpc.BidiStreamingServer[FileWatcherEvent, FileWatcherEvent]) error
	GetReusableToken(context.Context, *GetReusableTokenRequest) (*GetReusableTokenResponse, error)
	mustEmbedUnimplementedCodioCmdServer()
}

// UnimplementedCodioCmdServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCodioCmdServer struct{}

func (UnimplementedCodioCmdServer) FileUpload(grpc.ClientStreamingServer[FileUploadRequest, FileUploadResponse]) error {
	return status.Errorf(codes.Unimplemented, "method FileUpload not implemented")
}
func (UnimplementedCodioCmdServer) FileDownload(*FileDownloadRequest, grpc.ServerStreamingServer[FileDownloadResponse]) error {
	return status.Errorf(codes.Unimplemented, "method FileDownload not implemented")
}
func (UnimplementedCodioCmdServer) Exec(context.Context, *ExecRequest) (*ExecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exec not implemented")
}
func (UnimplementedCodioCmdServer) ExecAsync(context.Context, *ExecAsyncRequest) (*ExecAsyncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecAsync not implemented")
}
func (UnimplementedCodioCmdServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedCodioCmdServer) GetVMHostName(context.Context, *GetVMHostNameRequest) (*GetVMHostNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVMHostName not implemented")
}
func (UnimplementedCodioCmdServer) GetRemoteFileList(*GetRemoteFileListRequest, grpc.ServerStreamingServer[GetRemoteFileListResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetRemoteFileList not implemented")
}
func (UnimplementedCodioCmdServer) SyncPath(grpc.BidiStreamingServer[FileWatcherEvent, FileWatcherEvent]) error {
	return status.Errorf(codes.Unimplemented, "method SyncPath not implemented")
}
func (UnimplementedCodioCmdServer) GetReusableToken(context.Context, *GetReusableTokenRequest) (*GetReusableTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReusableToken not implemented")
}
func (UnimplementedCodioCmdServer) mustEmbedUnimplementedCodioCmdServer() {}
func (UnimplementedCodioCmdServer) testEmbeddedByValue()                  {}

// UnsafeCodioCmdServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CodioCmdServer will
// result in compilation errors.
type UnsafeCodioCmdServer interface {
	mustEmbedUnimplementedCodioCmdServer()
}

func RegisterCodioCmdServer(s grpc.ServiceRegistrar, srv CodioCmdServer) {
	// If the following call pancis, it indicates UnimplementedCodioCmdServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CodioCmd_ServiceDesc, srv)
}

func _CodioCmd_FileUpload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CodioCmdServer).FileUpload(&grpc.GenericServerStream[FileUploadRequest, FileUploadResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CodioCmd_FileUploadServer = grpc.ClientStreamingServer[FileUploadRequest, FileUploadResponse]

func _CodioCmd_FileDownload_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FileDownloadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CodioCmdServer).FileDownload(m, &grpc.GenericServerStream[FileDownloadRequest, FileDownloadResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CodioCmd_FileDownloadServer = grpc.ServerStreamingServer[FileDownloadResponse]

func _CodioCmd_Exec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodioCmdServer).Exec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodioCmd_Exec_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodioCmdServer).Exec(ctx, req.(*ExecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodioCmd_ExecAsync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecAsyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodioCmdServer).ExecAsync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodioCmd_ExecAsync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodioCmdServer).ExecAsync(ctx, req.(*ExecAsyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodioCmd_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodioCmdServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodioCmd_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodioCmdServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodioCmd_GetVMHostName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVMHostNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodioCmdServer).GetVMHostName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodioCmd_GetVMHostName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodioCmdServer).GetVMHostName(ctx, req.(*GetVMHostNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodioCmd_GetRemoteFileList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetRemoteFileListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CodioCmdServer).GetRemoteFileList(m, &grpc.GenericServerStream[GetRemoteFileListRequest, GetRemoteFileListResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CodioCmd_GetRemoteFileListServer = grpc.ServerStreamingServer[GetRemoteFileListResponse]

func _CodioCmd_SyncPath_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CodioCmdServer).SyncPath(&grpc.GenericServerStream[FileWatcherEvent, FileWatcherEvent]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CodioCmd_SyncPathServer = grpc.BidiStreamingServer[FileWatcherEvent, FileWatcherEvent]

func _CodioCmd_GetReusableToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReusableTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodioCmdServer).GetReusableToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodioCmd_GetReusableToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodioCmdServer).GetReusableToken(ctx, req.(*GetReusableTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CodioCmd_ServiceDesc is the grpc.ServiceDesc for CodioCmd service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CodioCmd_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.codio.chipmunk.proto.codiocmd.CodioCmd",
	HandlerType: (*CodioCmdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Exec",
			Handler:    _CodioCmd_Exec_Handler,
		},
		{
			MethodName: "ExecAsync",
			Handler:    _CodioCmd_ExecAsync_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _CodioCmd_Ping_Handler,
		},
		{
			MethodName: "GetVMHostName",
			Handler:    _CodioCmd_GetVMHostName_Handler,
		},
		{
			MethodName: "GetReusableToken",
			Handler:    _CodioCmd_GetReusableToken_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FileUpload",
			Handler:       _CodioCmd_FileUpload_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FileDownload",
			Handler:       _CodioCmd_FileDownload_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetRemoteFileList",
			Handler:       _CodioCmd_GetRemoteFileList_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SyncPath",
			Handler:       _CodioCmd_SyncPath_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/codiocmd/codiocmd_service.proto",
}
