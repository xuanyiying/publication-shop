// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: api/catalog/v1/catalog.proto

package v1

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
	Catalog_CreateCatalog_FullMethodName = "/api.catalog.v1.Catalog/CreateCatalog"
	Catalog_UpdateCatalog_FullMethodName = "/api.catalog.v1.Catalog/UpdateCatalog"
	Catalog_DeleteCatalog_FullMethodName = "/api.catalog.v1.Catalog/DeleteCatalog"
	Catalog_GetCatalog_FullMethodName    = "/api.catalog.v1.Catalog/GetCatalog"
	Catalog_ListCatalog_FullMethodName   = "/api.catalog.v1.Catalog/ListCatalog"
)

// CatalogClient is the client API for Catalog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatalogClient interface {
	CreateCatalog(ctx context.Context, in *CreateCatalogRequest, opts ...grpc.CallOption) (*CreateCatalogReply, error)
	UpdateCatalog(ctx context.Context, in *UpdateCatalogRequest, opts ...grpc.CallOption) (*UpdateCatalogReply, error)
	DeleteCatalog(ctx context.Context, in *DeleteCatalogRequest, opts ...grpc.CallOption) (*DeleteCatalogReply, error)
	GetCatalog(ctx context.Context, in *GetCatalogRequest, opts ...grpc.CallOption) (*GetCatalogReply, error)
	ListCatalog(ctx context.Context, in *ListCatalogRequest, opts ...grpc.CallOption) (*ListCatalogReply, error)
}

type catalogClient struct {
	cc grpc.ClientConnInterface
}

func NewCatalogClient(cc grpc.ClientConnInterface) CatalogClient {
	return &catalogClient{cc}
}

func (c *catalogClient) CreateCatalog(ctx context.Context, in *CreateCatalogRequest, opts ...grpc.CallOption) (*CreateCatalogReply, error) {
	out := new(CreateCatalogReply)
	err := c.cc.Invoke(ctx, Catalog_CreateCatalog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) UpdateCatalog(ctx context.Context, in *UpdateCatalogRequest, opts ...grpc.CallOption) (*UpdateCatalogReply, error) {
	out := new(UpdateCatalogReply)
	err := c.cc.Invoke(ctx, Catalog_UpdateCatalog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) DeleteCatalog(ctx context.Context, in *DeleteCatalogRequest, opts ...grpc.CallOption) (*DeleteCatalogReply, error) {
	out := new(DeleteCatalogReply)
	err := c.cc.Invoke(ctx, Catalog_DeleteCatalog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) GetCatalog(ctx context.Context, in *GetCatalogRequest, opts ...grpc.CallOption) (*GetCatalogReply, error) {
	out := new(GetCatalogReply)
	err := c.cc.Invoke(ctx, Catalog_GetCatalog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) ListCatalog(ctx context.Context, in *ListCatalogRequest, opts ...grpc.CallOption) (*ListCatalogReply, error) {
	out := new(ListCatalogReply)
	err := c.cc.Invoke(ctx, Catalog_ListCatalog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogServer is the server API for Catalog service.
// All implementations must embed UnimplementedCatalogServer
// for forward compatibility
type CatalogServer interface {
	CreateCatalog(context.Context, *CreateCatalogRequest) (*CreateCatalogReply, error)
	UpdateCatalog(context.Context, *UpdateCatalogRequest) (*UpdateCatalogReply, error)
	DeleteCatalog(context.Context, *DeleteCatalogRequest) (*DeleteCatalogReply, error)
	GetCatalog(context.Context, *GetCatalogRequest) (*GetCatalogReply, error)
	ListCatalog(context.Context, *ListCatalogRequest) (*ListCatalogReply, error)
	mustEmbedUnimplementedCatalogServer()
}

// UnimplementedCatalogServer must be embedded to have forward compatible implementations.
type UnimplementedCatalogServer struct {
}

func (UnimplementedCatalogServer) CreateCatalog(context.Context, *CreateCatalogRequest) (*CreateCatalogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCatalog not implemented")
}
func (UnimplementedCatalogServer) UpdateCatalog(context.Context, *UpdateCatalogRequest) (*UpdateCatalogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCatalog not implemented")
}
func (UnimplementedCatalogServer) DeleteCatalog(context.Context, *DeleteCatalogRequest) (*DeleteCatalogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCatalog not implemented")
}
func (UnimplementedCatalogServer) GetCatalog(context.Context, *GetCatalogRequest) (*GetCatalogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCatalog not implemented")
}
func (UnimplementedCatalogServer) ListCatalog(context.Context, *ListCatalogRequest) (*ListCatalogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCatalog not implemented")
}
func (UnimplementedCatalogServer) mustEmbedUnimplementedCatalogServer() {}

// UnsafeCatalogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatalogServer will
// result in compilation errors.
type UnsafeCatalogServer interface {
	mustEmbedUnimplementedCatalogServer()
}

func RegisterCatalogServer(s grpc.ServiceRegistrar, srv CatalogServer) {
	s.RegisterService(&Catalog_ServiceDesc, srv)
}

func _Catalog_CreateCatalog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCatalogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).CreateCatalog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Catalog_CreateCatalog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).CreateCatalog(ctx, req.(*CreateCatalogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_UpdateCatalog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCatalogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).UpdateCatalog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Catalog_UpdateCatalog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).UpdateCatalog(ctx, req.(*UpdateCatalogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_DeleteCatalog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCatalogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).DeleteCatalog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Catalog_DeleteCatalog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).DeleteCatalog(ctx, req.(*DeleteCatalogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_GetCatalog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCatalogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).GetCatalog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Catalog_GetCatalog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).GetCatalog(ctx, req.(*GetCatalogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_ListCatalog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCatalogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).ListCatalog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Catalog_ListCatalog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).ListCatalog(ctx, req.(*ListCatalogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Catalog_ServiceDesc is the grpc.ServiceDesc for Catalog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Catalog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.catalog.v1.Catalog",
	HandlerType: (*CatalogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCatalog",
			Handler:    _Catalog_CreateCatalog_Handler,
		},
		{
			MethodName: "UpdateCatalog",
			Handler:    _Catalog_UpdateCatalog_Handler,
		},
		{
			MethodName: "DeleteCatalog",
			Handler:    _Catalog_DeleteCatalog_Handler,
		},
		{
			MethodName: "GetCatalog",
			Handler:    _Catalog_GetCatalog_Handler,
		},
		{
			MethodName: "ListCatalog",
			Handler:    _Catalog_ListCatalog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/catalog/v1/catalog.proto",
}
