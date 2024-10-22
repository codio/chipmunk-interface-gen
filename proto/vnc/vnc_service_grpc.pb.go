// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.23.4
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
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Vnc_GetStatus_FullMethodName          = "/com.codio.chipmunk.proto.vnc.Vnc/GetStatus"
	Vnc_FetchVncPassword_FullMethodName   = "/com.codio.chipmunk.proto.vnc.Vnc/FetchVncPassword"
	Vnc_RestartVncService_FullMethodName  = "/com.codio.chipmunk.proto.vnc.Vnc/RestartVncService"
	Vnc_ForceUpdateConfig_FullMethodName  = "/com.codio.chipmunk.proto.vnc.Vnc/ForceUpdateConfig"
	Vnc_GetChipmunkVersion_FullMethodName = "/com.codio.chipmunk.proto.vnc.Vnc/GetChipmunkVersion"
)

// VncClient is the client API for Vnc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VncClient interface {
	GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error)
	FetchVncPassword(ctx context.Context, in *FetchVncPasswordRequest, opts ...grpc.CallOption) (*FetchVncPasswordResponse, error)
	RestartVncService(ctx context.Context, in *RestartVncServiceRequest, opts ...grpc.CallOption) (*RestartVncServiceResponse, error)
	ForceUpdateConfig(ctx context.Context, in *ForceUpdateConfigRequest, opts ...grpc.CallOption) (*ForceUpdateConfigResponse, error)
	GetChipmunkVersion(ctx context.Context, in *GetChipmunkVersionRequest, opts ...grpc.CallOption) (*GetChipmunkVersionResponse, error)
}

type vncClient struct {
	cc grpc.ClientConnInterface
}

func NewVncClient(cc grpc.ClientConnInterface) VncClient {
	return &vncClient{cc}
}

func (c *vncClient) GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStatusResponse)
	err := c.cc.Invoke(ctx, Vnc_GetStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vncClient) FetchVncPassword(ctx context.Context, in *FetchVncPasswordRequest, opts ...grpc.CallOption) (*FetchVncPasswordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FetchVncPasswordResponse)
	err := c.cc.Invoke(ctx, Vnc_FetchVncPassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vncClient) RestartVncService(ctx context.Context, in *RestartVncServiceRequest, opts ...grpc.CallOption) (*RestartVncServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RestartVncServiceResponse)
	err := c.cc.Invoke(ctx, Vnc_RestartVncService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vncClient) ForceUpdateConfig(ctx context.Context, in *ForceUpdateConfigRequest, opts ...grpc.CallOption) (*ForceUpdateConfigResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ForceUpdateConfigResponse)
	err := c.cc.Invoke(ctx, Vnc_ForceUpdateConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vncClient) GetChipmunkVersion(ctx context.Context, in *GetChipmunkVersionRequest, opts ...grpc.CallOption) (*GetChipmunkVersionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetChipmunkVersionResponse)
	err := c.cc.Invoke(ctx, Vnc_GetChipmunkVersion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VncServer is the server API for Vnc service.
// All implementations must embed UnimplementedVncServer
// for forward compatibility.
type VncServer interface {
	GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error)
	FetchVncPassword(context.Context, *FetchVncPasswordRequest) (*FetchVncPasswordResponse, error)
	RestartVncService(context.Context, *RestartVncServiceRequest) (*RestartVncServiceResponse, error)
	ForceUpdateConfig(context.Context, *ForceUpdateConfigRequest) (*ForceUpdateConfigResponse, error)
	GetChipmunkVersion(context.Context, *GetChipmunkVersionRequest) (*GetChipmunkVersionResponse, error)
	mustEmbedUnimplementedVncServer()
}

// UnimplementedVncServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedVncServer struct{}

func (UnimplementedVncServer) GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedVncServer) FetchVncPassword(context.Context, *FetchVncPasswordRequest) (*FetchVncPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchVncPassword not implemented")
}
func (UnimplementedVncServer) RestartVncService(context.Context, *RestartVncServiceRequest) (*RestartVncServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestartVncService not implemented")
}
func (UnimplementedVncServer) ForceUpdateConfig(context.Context, *ForceUpdateConfigRequest) (*ForceUpdateConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForceUpdateConfig not implemented")
}
func (UnimplementedVncServer) GetChipmunkVersion(context.Context, *GetChipmunkVersionRequest) (*GetChipmunkVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChipmunkVersion not implemented")
}
func (UnimplementedVncServer) mustEmbedUnimplementedVncServer() {}
func (UnimplementedVncServer) testEmbeddedByValue()             {}

// UnsafeVncServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VncServer will
// result in compilation errors.
type UnsafeVncServer interface {
	mustEmbedUnimplementedVncServer()
}

func RegisterVncServer(s grpc.ServiceRegistrar, srv VncServer) {
	// If the following call pancis, it indicates UnimplementedVncServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
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
		FullMethod: Vnc_GetStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VncServer).GetStatus(ctx, req.(*GetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vnc_FetchVncPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchVncPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VncServer).FetchVncPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Vnc_FetchVncPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VncServer).FetchVncPassword(ctx, req.(*FetchVncPasswordRequest))
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
		FullMethod: Vnc_RestartVncService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VncServer).RestartVncService(ctx, req.(*RestartVncServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vnc_ForceUpdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForceUpdateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VncServer).ForceUpdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Vnc_ForceUpdateConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VncServer).ForceUpdateConfig(ctx, req.(*ForceUpdateConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vnc_GetChipmunkVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChipmunkVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VncServer).GetChipmunkVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Vnc_GetChipmunkVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VncServer).GetChipmunkVersion(ctx, req.(*GetChipmunkVersionRequest))
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
			MethodName: "FetchVncPassword",
			Handler:    _Vnc_FetchVncPassword_Handler,
		},
		{
			MethodName: "RestartVncService",
			Handler:    _Vnc_RestartVncService_Handler,
		},
		{
			MethodName: "ForceUpdateConfig",
			Handler:    _Vnc_ForceUpdateConfig_Handler,
		},
		{
			MethodName: "GetChipmunkVersion",
			Handler:    _Vnc_GetChipmunkVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/vnc/vnc_service.proto",
}
