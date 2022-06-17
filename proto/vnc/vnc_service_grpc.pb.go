// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: proto/vnc/vnc_service.proto

package vnc

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

// VncClient is the client API for Vnc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VncClient interface {
	GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error)
	CreateVncPassword(ctx context.Context, in *CreateVncPasswordRequest, opts ...grpc.CallOption) (*CreateVncPasswordResponse, error)
	RestartVncService(ctx context.Context, in *RestartVncServiceRequest, opts ...grpc.CallOption) (*RestartVncServiceResponse, error)
	SetChipmunkVersion(ctx context.Context, in *SetChipmunkVersionRequest, opts ...grpc.CallOption) (*SetChipmunkVersionResponse, error)
}

type vncClient struct {
	cc grpc.ClientConnInterface
}

func NewVncClient(cc grpc.ClientConnInterface) VncClient {
	return &vncClient{cc}
}

func (c *vncClient) GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error) {
	out := new(GetStatusResponse)
	err := c.cc.Invoke(ctx, "/com.codio.chipmunk.proto.vnc.Vnc/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vncClient) CreateVncPassword(ctx context.Context, in *CreateVncPasswordRequest, opts ...grpc.CallOption) (*CreateVncPasswordResponse, error) {
	out := new(CreateVncPasswordResponse)
	err := c.cc.Invoke(ctx, "/com.codio.chipmunk.proto.vnc.Vnc/CreateVncPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vncClient) RestartVncService(ctx context.Context, in *RestartVncServiceRequest, opts ...grpc.CallOption) (*RestartVncServiceResponse, error) {
	out := new(RestartVncServiceResponse)
	err := c.cc.Invoke(ctx, "/com.codio.chipmunk.proto.vnc.Vnc/RestartVncService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vncClient) SetChipmunkVersion(ctx context.Context, in *SetChipmunkVersionRequest, opts ...grpc.CallOption) (*SetChipmunkVersionResponse, error) {
	out := new(SetChipmunkVersionResponse)
	err := c.cc.Invoke(ctx, "/com.codio.chipmunk.proto.vnc.Vnc/SetChipmunkVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VncServer is the server API for Vnc service.
// All implementations must embed UnimplementedVncServer
// for forward compatibility
type VncServer interface {
	GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error)
	CreateVncPassword(context.Context, *CreateVncPasswordRequest) (*CreateVncPasswordResponse, error)
	RestartVncService(context.Context, *RestartVncServiceRequest) (*RestartVncServiceResponse, error)
	SetChipmunkVersion(context.Context, *SetChipmunkVersionRequest) (*SetChipmunkVersionResponse, error)
	mustEmbedUnimplementedVncServer()
}

// UnimplementedVncServer must be embedded to have forward compatible implementations.
type UnimplementedVncServer struct {
}

func (UnimplementedVncServer) GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedVncServer) CreateVncPassword(context.Context, *CreateVncPasswordRequest) (*CreateVncPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVncPassword not implemented")
}
func (UnimplementedVncServer) RestartVncService(context.Context, *RestartVncServiceRequest) (*RestartVncServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestartVncService not implemented")
}
func (UnimplementedVncServer) SetChipmunkVersion(context.Context, *SetChipmunkVersionRequest) (*SetChipmunkVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetChipmunkVersion not implemented")
}
func (UnimplementedVncServer) mustEmbedUnimplementedVncServer() {}

// UnsafeVncServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VncServer will
// result in compilation errors.
type UnsafeVncServer interface {
	mustEmbedUnimplementedVncServer()
}

func RegisterVncServer(s grpc.ServiceRegistrar, srv VncServer) {
	s.RegisterService(&Vnc_ServiceDesc, srv)
}

func _Vnc_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VncServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.codio.chipmunk.proto.vnc.Vnc/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VncServer).GetStatus(ctx, req.(*GetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vnc_CreateVncPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVncPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VncServer).CreateVncPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.codio.chipmunk.proto.vnc.Vnc/CreateVncPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VncServer).CreateVncPassword(ctx, req.(*CreateVncPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vnc_RestartVncService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestartVncServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VncServer).RestartVncService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.codio.chipmunk.proto.vnc.Vnc/RestartVncService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VncServer).RestartVncService(ctx, req.(*RestartVncServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vnc_SetChipmunkVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetChipmunkVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VncServer).SetChipmunkVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.codio.chipmunk.proto.vnc.Vnc/SetChipmunkVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VncServer).SetChipmunkVersion(ctx, req.(*SetChipmunkVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Vnc_ServiceDesc is the grpc.ServiceDesc for Vnc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Vnc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.codio.chipmunk.proto.vnc.Vnc",
	HandlerType: (*VncServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _Vnc_GetStatus_Handler,
		},
		{
			MethodName: "CreateVncPassword",
			Handler:    _Vnc_CreateVncPassword_Handler,
		},
		{
			MethodName: "RestartVncService",
			Handler:    _Vnc_RestartVncService_Handler,
		},
		{
			MethodName: "SetChipmunkVersion",
			Handler:    _Vnc_SetChipmunkVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/vnc/vnc_service.proto",
}
