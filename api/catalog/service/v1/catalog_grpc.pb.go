// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.11.4
// source: api/catalog/service/v1/catalog.proto

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

// CatalogClient is the client API for Catalog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatalogClient interface {
	ListPublication(ctx context.Context, in *ListPublicationReq, opts ...grpc.CallOption) (*ListPublicationReply, error)
	ListPublicationNextToken(ctx context.Context, in *ListPublicationNextTokenReq, opts ...grpc.CallOption) (*ListPublicationReplyNextToken, error)
	CreatePublication(ctx context.Context, in *CreatePublicationReq, opts ...grpc.CallOption) (*CreatePublicationReply, error)
	GetPublication(ctx context.Context, in *GetPublicationReq, opts ...grpc.CallOption) (*GetPublicationReply, error)
	UpdatePublication(ctx context.Context, in *UpdatePublicationReq, opts ...grpc.CallOption) (*UpdatePublicationReply, error)
	DeletePublication(ctx context.Context, in *DeletePublicationReq, opts ...grpc.CallOption) (*DeletePublicationReply, error)
}

type catalogClient struct {
	cc grpc.ClientConnInterface
}

func NewCatalogClient(cc grpc.ClientConnInterface) CatalogClient {
	return &catalogClient{cc}
}

func (c *catalogClient) ListPublication(ctx context.Context, in *ListPublicationReq, opts ...grpc.CallOption) (*ListPublicationReply, error) {
	out := new(ListPublicationReply)
	err := c.cc.Invoke(ctx, "/catalog.service.v1.Catalog/ListPublication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) ListPublicationNextToken(ctx context.Context, in *ListPublicationNextTokenReq, opts ...grpc.CallOption) (*ListPublicationReplyNextToken, error) {
	out := new(ListPublicationReplyNextToken)
	err := c.cc.Invoke(ctx, "/catalog.service.v1.Catalog/ListPublicationNextToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) CreatePublication(ctx context.Context, in *CreatePublicationReq, opts ...grpc.CallOption) (*CreatePublicationReply, error) {
	out := new(CreatePublicationReply)
	err := c.cc.Invoke(ctx, "/catalog.service.v1.Catalog/CreatePublication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) GetPublication(ctx context.Context, in *GetPublicationReq, opts ...grpc.CallOption) (*GetPublicationReply, error) {
	out := new(GetPublicationReply)
	err := c.cc.Invoke(ctx, "/catalog.service.v1.Catalog/GetPublication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) UpdatePublication(ctx context.Context, in *UpdatePublicationReq, opts ...grpc.CallOption) (*UpdatePublicationReply, error) {
	out := new(UpdatePublicationReply)
	err := c.cc.Invoke(ctx, "/catalog.service.v1.Catalog/UpdatePublication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) DeletePublication(ctx context.Context, in *DeletePublicationReq, opts ...grpc.CallOption) (*DeletePublicationReply, error) {
	out := new(DeletePublicationReply)
	err := c.cc.Invoke(ctx, "/catalog.service.v1.Catalog/DeletePublication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogServer is the server API for Catalog service.
// All implementations must embed UnimplementedCatalogServer
// for forward compatibility
type CatalogServer interface {
	ListPublication(context.Context, *ListPublicationReq) (*ListPublicationReply, error)
	ListPublicationNextToken(context.Context, *ListPublicationNextTokenReq) (*ListPublicationReplyNextToken, error)
	CreatePublication(context.Context, *CreatePublicationReq) (*CreatePublicationReply, error)
	GetPublication(context.Context, *GetPublicationReq) (*GetPublicationReply, error)
	UpdatePublication(context.Context, *UpdatePublicationReq) (*UpdatePublicationReply, error)
	DeletePublication(context.Context, *DeletePublicationReq) (*DeletePublicationReply, error)
	mustEmbedUnimplementedCatalogServer()
}

// UnimplementedCatalogServer must be embedded to have forward compatible implementations.
type UnimplementedCatalogServer struct {
}

func (UnimplementedCatalogServer) ListPublication(context.Context, *ListPublicationReq) (*ListPublicationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPublication not implemented")
}
func (UnimplementedCatalogServer) ListPublicationNextToken(context.Context, *ListPublicationNextTokenReq) (*ListPublicationReplyNextToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPublicationNextToken not implemented")
}
func (UnimplementedCatalogServer) CreatePublication(context.Context, *CreatePublicationReq) (*CreatePublicationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePublication not implemented")
}
func (UnimplementedCatalogServer) GetPublication(context.Context, *GetPublicationReq) (*GetPublicationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublication not implemented")
}
func (UnimplementedCatalogServer) UpdatePublication(context.Context, *UpdatePublicationReq) (*UpdatePublicationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePublication not implemented")
}
func (UnimplementedCatalogServer) DeletePublication(context.Context, *DeletePublicationReq) (*DeletePublicationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePublication not implemented")
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

func _Catalog_ListPublication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPublicationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).ListPublication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.service.v1.Catalog/ListPublication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).ListPublication(ctx, req.(*ListPublicationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_ListPublicationNextToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPublicationNextTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).ListPublicationNextToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.service.v1.Catalog/ListPublicationNextToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).ListPublicationNextToken(ctx, req.(*ListPublicationNextTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_CreatePublication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePublicationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).CreatePublication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.service.v1.Catalog/CreatePublication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).CreatePublication(ctx, req.(*CreatePublicationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_GetPublication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).GetPublication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.service.v1.Catalog/GetPublication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).GetPublication(ctx, req.(*GetPublicationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_UpdatePublication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePublicationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).UpdatePublication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.service.v1.Catalog/UpdatePublication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).UpdatePublication(ctx, req.(*UpdatePublicationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_DeletePublication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePublicationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).DeletePublication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.service.v1.Catalog/DeletePublication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).DeletePublication(ctx, req.(*DeletePublicationReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Catalog_ServiceDesc is the grpc.ServiceDesc for Catalog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Catalog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "catalog.service.v1.Catalog",
	HandlerType: (*CatalogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPublication",
			Handler:    _Catalog_ListPublication_Handler,
		},
		{
			MethodName: "ListPublicationNextToken",
			Handler:    _Catalog_ListPublicationNextToken_Handler,
		},
		{
			MethodName: "CreatePublication",
			Handler:    _Catalog_CreatePublication_Handler,
		},
		{
			MethodName: "GetPublication",
			Handler:    _Catalog_GetPublication_Handler,
		},
		{
			MethodName: "UpdatePublication",
			Handler:    _Catalog_UpdatePublication_Handler,
		},
		{
			MethodName: "DeletePublication",
			Handler:    _Catalog_DeletePublication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/catalog/service/v1/catalog.proto",
}
