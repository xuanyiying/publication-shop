// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: api/github.com/xuanyiying/publication-shop/app/transaction/service/v1/transaction.proto

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
	Transaction_CreateTx_FullMethodName            = "/api.transaction.service.v1.Transaction/CreateTx"
	Transaction_UpdateTx_FullMethodName            = "/api.transaction.service.v1.Transaction/UpdateTx"
	Transaction_DeleteTx_FullMethodName            = "/api.transaction.service.v1.Transaction/DeleteTx"
	Transaction_GetTx_FullMethodName               = "/api.transaction.service.v1.Transaction/GetTx"
	Transaction_ListTx_FullMethodName              = "/api.transaction.service.v1.Transaction/ListTx"
	Transaction_GetTxByTxNo_FullMethodName         = "/api.transaction.service.v1.Transaction/GetTxByTxNo"
	Transaction_ListTxByTxType_FullMethodName      = "/api.transaction.service.v1.Transaction/ListTxByTxType"
	Transaction_ListTxByUserId_FullMethodName      = "/api.transaction.service.v1.Transaction/ListTxByUserId"
	Transaction_ListTxByPaymentId_FullMethodName   = "/api.transaction.service.v1.Transaction/ListTxByPaymentId"
	Transaction_ListTxByTxStatus_FullMethodName    = "/api.transaction.service.v1.Transaction/ListTxByTxStatus"
	Transaction_ListTxByTxDateRange_FullMethodName = "/api.transaction.service.v1.Transaction/ListTxByTxDateRange"
)

// TransactionClient is the client API for Transaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionClient interface {
	CreateTx(ctx context.Context, in *CreateTxReq, opts ...grpc.CallOption) (*CreateTxResp, error)
	UpdateTx(ctx context.Context, in *UpdateTxReq, opts ...grpc.CallOption) (*UpdateTxResp, error)
	DeleteTx(ctx context.Context, in *DeleteTxReq, opts ...grpc.CallOption) (*DeleteTxResp, error)
	GetTx(ctx context.Context, in *GetTxReq, opts ...grpc.CallOption) (*TxResp, error)
	ListTx(ctx context.Context, in *ListTxReq, opts ...grpc.CallOption) (*ListTxResp, error)
	GetTxByTxNo(ctx context.Context, in *GetTxReq, opts ...grpc.CallOption) (*TxResp, error)
	ListTxByTxType(ctx context.Context, in *ListTxReq, opts ...grpc.CallOption) (*ListTxResp, error)
	ListTxByUserId(ctx context.Context, in *ListTxReq, opts ...grpc.CallOption) (*ListTxResp, error)
	ListTxByPaymentId(ctx context.Context, in *ListTxReq, opts ...grpc.CallOption) (*ListTxResp, error)
	ListTxByTxStatus(ctx context.Context, in *ListTxReq, opts ...grpc.CallOption) (*ListTxResp, error)
	ListTxByTxDateRange(ctx context.Context, in *ListTxDateRangeReq, opts ...grpc.CallOption) (*ListTxResp, error)
}

type transactionClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionClient(cc grpc.ClientConnInterface) TransactionClient {
	return &transactionClient{cc}
}

func (c *transactionClient) CreateTx(ctx context.Context, in *CreateTxReq, opts ...grpc.CallOption) (*CreateTxResp, error) {
	out := new(CreateTxResp)
	err := c.cc.Invoke(ctx, Transaction_CreateTx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) UpdateTx(ctx context.Context, in *UpdateTxReq, opts ...grpc.CallOption) (*UpdateTxResp, error) {
	out := new(UpdateTxResp)
	err := c.cc.Invoke(ctx, Transaction_UpdateTx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) DeleteTx(ctx context.Context, in *DeleteTxReq, opts ...grpc.CallOption) (*DeleteTxResp, error) {
	out := new(DeleteTxResp)
	err := c.cc.Invoke(ctx, Transaction_DeleteTx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) GetTx(ctx context.Context, in *GetTxReq, opts ...grpc.CallOption) (*TxResp, error) {
	out := new(TxResp)
	err := c.cc.Invoke(ctx, Transaction_GetTx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) ListTx(ctx context.Context, in *ListTxReq, opts ...grpc.CallOption) (*ListTxResp, error) {
	out := new(ListTxResp)
	err := c.cc.Invoke(ctx, Transaction_ListTx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) GetTxByTxNo(ctx context.Context, in *GetTxReq, opts ...grpc.CallOption) (*TxResp, error) {
	out := new(TxResp)
	err := c.cc.Invoke(ctx, Transaction_GetTxByTxNo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) ListTxByTxType(ctx context.Context, in *ListTxReq, opts ...grpc.CallOption) (*ListTxResp, error) {
	out := new(ListTxResp)
	err := c.cc.Invoke(ctx, Transaction_ListTxByTxType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) ListTxByUserId(ctx context.Context, in *ListTxReq, opts ...grpc.CallOption) (*ListTxResp, error) {
	out := new(ListTxResp)
	err := c.cc.Invoke(ctx, Transaction_ListTxByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) ListTxByPaymentId(ctx context.Context, in *ListTxReq, opts ...grpc.CallOption) (*ListTxResp, error) {
	out := new(ListTxResp)
	err := c.cc.Invoke(ctx, Transaction_ListTxByPaymentId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) ListTxByTxStatus(ctx context.Context, in *ListTxReq, opts ...grpc.CallOption) (*ListTxResp, error) {
	out := new(ListTxResp)
	err := c.cc.Invoke(ctx, Transaction_ListTxByTxStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) ListTxByTxDateRange(ctx context.Context, in *ListTxDateRangeReq, opts ...grpc.CallOption) (*ListTxResp, error) {
	out := new(ListTxResp)
	err := c.cc.Invoke(ctx, Transaction_ListTxByTxDateRange_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServer is the server API for Transaction service.
// All implementations must embed UnimplementedTransactionServer
// for forward compatibility
type TransactionServer interface {
	CreateTx(context.Context, *CreateTxReq) (*CreateTxResp, error)
	UpdateTx(context.Context, *UpdateTxReq) (*UpdateTxResp, error)
	DeleteTx(context.Context, *DeleteTxReq) (*DeleteTxResp, error)
	GetTx(context.Context, *GetTxReq) (*TxResp, error)
	ListTx(context.Context, *ListTxReq) (*ListTxResp, error)
	GetTxByTxNo(context.Context, *GetTxReq) (*TxResp, error)
	ListTxByTxType(context.Context, *ListTxReq) (*ListTxResp, error)
	ListTxByUserId(context.Context, *ListTxReq) (*ListTxResp, error)
	ListTxByPaymentId(context.Context, *ListTxReq) (*ListTxResp, error)
	ListTxByTxStatus(context.Context, *ListTxReq) (*ListTxResp, error)
	ListTxByTxDateRange(context.Context, *ListTxDateRangeReq) (*ListTxResp, error)
	mustEmbedUnimplementedTransactionServer()
}

// UnimplementedTransactionServer must be embedded to have forward compatible implementations.
type UnimplementedTransactionServer struct {
}

func (UnimplementedTransactionServer) CreateTx(context.Context, *CreateTxReq) (*CreateTxResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTx not implemented")
}
func (UnimplementedTransactionServer) UpdateTx(context.Context, *UpdateTxReq) (*UpdateTxResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTx not implemented")
}
func (UnimplementedTransactionServer) DeleteTx(context.Context, *DeleteTxReq) (*DeleteTxResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTx not implemented")
}
func (UnimplementedTransactionServer) GetTx(context.Context, *GetTxReq) (*TxResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTx not implemented")
}
func (UnimplementedTransactionServer) ListTx(context.Context, *ListTxReq) (*ListTxResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTx not implemented")
}
func (UnimplementedTransactionServer) GetTxByTxNo(context.Context, *GetTxReq) (*TxResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxByTxNo not implemented")
}
func (UnimplementedTransactionServer) ListTxByTxType(context.Context, *ListTxReq) (*ListTxResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTxByTxType not implemented")
}
func (UnimplementedTransactionServer) ListTxByUserId(context.Context, *ListTxReq) (*ListTxResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTxByUserId not implemented")
}
func (UnimplementedTransactionServer) ListTxByPaymentId(context.Context, *ListTxReq) (*ListTxResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTxByPaymentId not implemented")
}
func (UnimplementedTransactionServer) ListTxByTxStatus(context.Context, *ListTxReq) (*ListTxResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTxByTxStatus not implemented")
}
func (UnimplementedTransactionServer) ListTxByTxDateRange(context.Context, *ListTxDateRangeReq) (*ListTxResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTxByTxDateRange not implemented")
}
func (UnimplementedTransactionServer) mustEmbedUnimplementedTransactionServer() {}

// UnsafeTransactionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionServer will
// result in compilation errors.
type UnsafeTransactionServer interface {
	mustEmbedUnimplementedTransactionServer()
}

func RegisterTransactionServer(s grpc.ServiceRegistrar, srv TransactionServer) {
	s.RegisterService(&Transaction_ServiceDesc, srv)
}

func _Transaction_CreateTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTxReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).CreateTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Transaction_CreateTx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).CreateTx(ctx, req.(*CreateTxReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_UpdateTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTxReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).UpdateTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Transaction_UpdateTx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).UpdateTx(ctx, req.(*UpdateTxReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_DeleteTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTxReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).DeleteTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Transaction_DeleteTx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).DeleteTx(ctx, req.(*DeleteTxReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_GetTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).GetTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Transaction_GetTx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).GetTx(ctx, req.(*GetTxReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_ListTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTxReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).ListTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Transaction_ListTx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).ListTx(ctx, req.(*ListTxReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_GetTxByTxNo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).GetTxByTxNo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Transaction_GetTxByTxNo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).GetTxByTxNo(ctx, req.(*GetTxReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_ListTxByTxType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTxReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).ListTxByTxType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Transaction_ListTxByTxType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).ListTxByTxType(ctx, req.(*ListTxReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_ListTxByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTxReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).ListTxByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Transaction_ListTxByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).ListTxByUserId(ctx, req.(*ListTxReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_ListTxByPaymentId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTxReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).ListTxByPaymentId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Transaction_ListTxByPaymentId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).ListTxByPaymentId(ctx, req.(*ListTxReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_ListTxByTxStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTxReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).ListTxByTxStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Transaction_ListTxByTxStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).ListTxByTxStatus(ctx, req.(*ListTxReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_ListTxByTxDateRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTxDateRangeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).ListTxByTxDateRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Transaction_ListTxByTxDateRange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).ListTxByTxDateRange(ctx, req.(*ListTxDateRangeReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Transaction_ServiceDesc is the grpc.ServiceDesc for Transaction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Transaction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.transaction.service.v1.Transaction",
	HandlerType: (*TransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTx",
			Handler:    _Transaction_CreateTx_Handler,
		},
		{
			MethodName: "UpdateTx",
			Handler:    _Transaction_UpdateTx_Handler,
		},
		{
			MethodName: "DeleteTx",
			Handler:    _Transaction_DeleteTx_Handler,
		},
		{
			MethodName: "GetTx",
			Handler:    _Transaction_GetTx_Handler,
		},
		{
			MethodName: "ListTx",
			Handler:    _Transaction_ListTx_Handler,
		},
		{
			MethodName: "GetTxByTxNo",
			Handler:    _Transaction_GetTxByTxNo_Handler,
		},
		{
			MethodName: "ListTxByTxType",
			Handler:    _Transaction_ListTxByTxType_Handler,
		},
		{
			MethodName: "ListTxByUserId",
			Handler:    _Transaction_ListTxByUserId_Handler,
		},
		{
			MethodName: "ListTxByPaymentId",
			Handler:    _Transaction_ListTxByPaymentId_Handler,
		},
		{
			MethodName: "ListTxByTxStatus",
			Handler:    _Transaction_ListTxByTxStatus_Handler,
		},
		{
			MethodName: "ListTxByTxDateRange",
			Handler:    _Transaction_ListTxByTxDateRange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/github.com/xuanyiying/publication-shop/app/transaction/service/v1/transaction.proto",
}
