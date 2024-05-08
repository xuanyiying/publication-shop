// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v5.26.1
// source: api/github.com/xuanyiying/publication-shop/app/transaction/service/v1/transaction.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationTransactionCreateTx = "/api.transaction.service.v1.Transaction/CreateTx"
const OperationTransactionDeleteTx = "/api.transaction.service.v1.Transaction/DeleteTx"
const OperationTransactionGetTx = "/api.transaction.service.v1.Transaction/GetTx"
const OperationTransactionGetTxByTxNo = "/api.transaction.service.v1.Transaction/GetTxByTxNo"
const OperationTransactionListTx = "/api.transaction.service.v1.Transaction/ListTx"
const OperationTransactionListTxByPaymentId = "/api.transaction.service.v1.Transaction/ListTxByPaymentId"
const OperationTransactionListTxByTxDateRange = "/api.transaction.service.v1.Transaction/ListTxByTxDateRange"
const OperationTransactionListTxByTxStatus = "/api.transaction.service.v1.Transaction/ListTxByTxStatus"
const OperationTransactionListTxByTxType = "/api.transaction.service.v1.Transaction/ListTxByTxType"
const OperationTransactionListTxByUserId = "/api.transaction.service.v1.Transaction/ListTxByUserId"
const OperationTransactionUpdateTx = "/api.transaction.service.v1.Transaction/UpdateTx"

type TransactionHTTPServer interface {
	CreateTx(context.Context, *CreateTxReq) (*CreateTxResp, error)
	DeleteTx(context.Context, *DeleteTxReq) (*DeleteTxResp, error)
	GetTx(context.Context, *GetTxReq) (*TxResp, error)
	GetTxByTxNo(context.Context, *GetTxReq) (*TxResp, error)
	ListTx(context.Context, *ListTxReq) (*ListTxResp, error)
	ListTxByPaymentId(context.Context, *ListTxReq) (*ListTxResp, error)
	ListTxByTxDateRange(context.Context, *ListTxDateRangeReq) (*ListTxResp, error)
	ListTxByTxStatus(context.Context, *ListTxReq) (*ListTxResp, error)
	ListTxByTxType(context.Context, *ListTxReq) (*ListTxResp, error)
	ListTxByUserId(context.Context, *ListTxReq) (*ListTxResp, error)
	UpdateTx(context.Context, *UpdateTxReq) (*UpdateTxResp, error)
}

func RegisterTransactionHTTPServer(s *http.Server, srv TransactionHTTPServer) {
	r := s.Route("/")
	r.POST("/api.transaction.service.v1.Transaction/CreateTx", _Transaction_CreateTx0_HTTP_Handler(srv))
	r.POST("/api.transaction.service.v1.Transaction/UpdateTx", _Transaction_UpdateTx0_HTTP_Handler(srv))
	r.POST("/api.transaction.service.v1.Transaction/DeleteTx", _Transaction_DeleteTx0_HTTP_Handler(srv))
	r.POST("/api.transaction.service.v1.Transaction/GetTx", _Transaction_GetTx0_HTTP_Handler(srv))
	r.POST("/api.transaction.service.v1.Transaction/ListTx", _Transaction_ListTx0_HTTP_Handler(srv))
	r.POST("/api.transaction.service.v1.Transaction/GetTxByTxNo", _Transaction_GetTxByTxNo0_HTTP_Handler(srv))
	r.POST("/api.transaction.service.v1.Transaction/ListTxByTxType", _Transaction_ListTxByTxType0_HTTP_Handler(srv))
	r.POST("/api.transaction.service.v1.Transaction/ListTxByUserId", _Transaction_ListTxByUserId0_HTTP_Handler(srv))
	r.POST("/api.transaction.service.v1.Transaction/ListTxByPaymentId", _Transaction_ListTxByPaymentId0_HTTP_Handler(srv))
	r.POST("/api.transaction.service.v1.Transaction/ListTxByTxStatus", _Transaction_ListTxByTxStatus0_HTTP_Handler(srv))
	r.POST("/api.transaction.service.v1.Transaction/ListTxByTxDateRange", _Transaction_ListTxByTxDateRange0_HTTP_Handler(srv))
}

func _Transaction_CreateTx0_HTTP_Handler(srv TransactionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateTxReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTransactionCreateTx)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateTx(ctx, req.(*CreateTxReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateTxResp)
		return ctx.Result(200, reply)
	}
}

func _Transaction_UpdateTx0_HTTP_Handler(srv TransactionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateTxReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTransactionUpdateTx)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateTx(ctx, req.(*UpdateTxReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateTxResp)
		return ctx.Result(200, reply)
	}
}

func _Transaction_DeleteTx0_HTTP_Handler(srv TransactionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteTxReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTransactionDeleteTx)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteTx(ctx, req.(*DeleteTxReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteTxResp)
		return ctx.Result(200, reply)
	}
}

func _Transaction_GetTx0_HTTP_Handler(srv TransactionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTxReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTransactionGetTx)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTx(ctx, req.(*GetTxReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TxResp)
		return ctx.Result(200, reply)
	}
}

func _Transaction_ListTx0_HTTP_Handler(srv TransactionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTxReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTransactionListTx)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListTx(ctx, req.(*ListTxReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTxResp)
		return ctx.Result(200, reply)
	}
}

func _Transaction_GetTxByTxNo0_HTTP_Handler(srv TransactionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTxReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTransactionGetTxByTxNo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTxByTxNo(ctx, req.(*GetTxReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TxResp)
		return ctx.Result(200, reply)
	}
}

func _Transaction_ListTxByTxType0_HTTP_Handler(srv TransactionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTxReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTransactionListTxByTxType)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListTxByTxType(ctx, req.(*ListTxReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTxResp)
		return ctx.Result(200, reply)
	}
}

func _Transaction_ListTxByUserId0_HTTP_Handler(srv TransactionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTxReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTransactionListTxByUserId)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListTxByUserId(ctx, req.(*ListTxReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTxResp)
		return ctx.Result(200, reply)
	}
}

func _Transaction_ListTxByPaymentId0_HTTP_Handler(srv TransactionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTxReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTransactionListTxByPaymentId)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListTxByPaymentId(ctx, req.(*ListTxReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTxResp)
		return ctx.Result(200, reply)
	}
}

func _Transaction_ListTxByTxStatus0_HTTP_Handler(srv TransactionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTxReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTransactionListTxByTxStatus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListTxByTxStatus(ctx, req.(*ListTxReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTxResp)
		return ctx.Result(200, reply)
	}
}

func _Transaction_ListTxByTxDateRange0_HTTP_Handler(srv TransactionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTxDateRangeReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTransactionListTxByTxDateRange)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListTxByTxDateRange(ctx, req.(*ListTxDateRangeReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTxResp)
		return ctx.Result(200, reply)
	}
}

type TransactionHTTPClient interface {
	CreateTx(ctx context.Context, req *CreateTxReq, opts ...http.CallOption) (rsp *CreateTxResp, err error)
	DeleteTx(ctx context.Context, req *DeleteTxReq, opts ...http.CallOption) (rsp *DeleteTxResp, err error)
	GetTx(ctx context.Context, req *GetTxReq, opts ...http.CallOption) (rsp *TxResp, err error)
	GetTxByTxNo(ctx context.Context, req *GetTxReq, opts ...http.CallOption) (rsp *TxResp, err error)
	ListTx(ctx context.Context, req *ListTxReq, opts ...http.CallOption) (rsp *ListTxResp, err error)
	ListTxByPaymentId(ctx context.Context, req *ListTxReq, opts ...http.CallOption) (rsp *ListTxResp, err error)
	ListTxByTxDateRange(ctx context.Context, req *ListTxDateRangeReq, opts ...http.CallOption) (rsp *ListTxResp, err error)
	ListTxByTxStatus(ctx context.Context, req *ListTxReq, opts ...http.CallOption) (rsp *ListTxResp, err error)
	ListTxByTxType(ctx context.Context, req *ListTxReq, opts ...http.CallOption) (rsp *ListTxResp, err error)
	ListTxByUserId(ctx context.Context, req *ListTxReq, opts ...http.CallOption) (rsp *ListTxResp, err error)
	UpdateTx(ctx context.Context, req *UpdateTxReq, opts ...http.CallOption) (rsp *UpdateTxResp, err error)
}

type TransactionHTTPClientImpl struct {
	cc *http.Client
}

func NewTransactionHTTPClient(client *http.Client) TransactionHTTPClient {
	return &TransactionHTTPClientImpl{client}
}

func (c *TransactionHTTPClientImpl) CreateTx(ctx context.Context, in *CreateTxReq, opts ...http.CallOption) (*CreateTxResp, error) {
	var out CreateTxResp
	pattern := "/api.transaction.service.v1.Transaction/CreateTx"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTransactionCreateTx))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TransactionHTTPClientImpl) DeleteTx(ctx context.Context, in *DeleteTxReq, opts ...http.CallOption) (*DeleteTxResp, error) {
	var out DeleteTxResp
	pattern := "/api.transaction.service.v1.Transaction/DeleteTx"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTransactionDeleteTx))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TransactionHTTPClientImpl) GetTx(ctx context.Context, in *GetTxReq, opts ...http.CallOption) (*TxResp, error) {
	var out TxResp
	pattern := "/api.transaction.service.v1.Transaction/GetTx"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTransactionGetTx))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TransactionHTTPClientImpl) GetTxByTxNo(ctx context.Context, in *GetTxReq, opts ...http.CallOption) (*TxResp, error) {
	var out TxResp
	pattern := "/api.transaction.service.v1.Transaction/GetTxByTxNo"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTransactionGetTxByTxNo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TransactionHTTPClientImpl) ListTx(ctx context.Context, in *ListTxReq, opts ...http.CallOption) (*ListTxResp, error) {
	var out ListTxResp
	pattern := "/api.transaction.service.v1.Transaction/ListTx"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTransactionListTx))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TransactionHTTPClientImpl) ListTxByPaymentId(ctx context.Context, in *ListTxReq, opts ...http.CallOption) (*ListTxResp, error) {
	var out ListTxResp
	pattern := "/api.transaction.service.v1.Transaction/ListTxByPaymentId"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTransactionListTxByPaymentId))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TransactionHTTPClientImpl) ListTxByTxDateRange(ctx context.Context, in *ListTxDateRangeReq, opts ...http.CallOption) (*ListTxResp, error) {
	var out ListTxResp
	pattern := "/api.transaction.service.v1.Transaction/ListTxByTxDateRange"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTransactionListTxByTxDateRange))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TransactionHTTPClientImpl) ListTxByTxStatus(ctx context.Context, in *ListTxReq, opts ...http.CallOption) (*ListTxResp, error) {
	var out ListTxResp
	pattern := "/api.transaction.service.v1.Transaction/ListTxByTxStatus"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTransactionListTxByTxStatus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TransactionHTTPClientImpl) ListTxByTxType(ctx context.Context, in *ListTxReq, opts ...http.CallOption) (*ListTxResp, error) {
	var out ListTxResp
	pattern := "/api.transaction.service.v1.Transaction/ListTxByTxType"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTransactionListTxByTxType))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TransactionHTTPClientImpl) ListTxByUserId(ctx context.Context, in *ListTxReq, opts ...http.CallOption) (*ListTxResp, error) {
	var out ListTxResp
	pattern := "/api.transaction.service.v1.Transaction/ListTxByUserId"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTransactionListTxByUserId))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TransactionHTTPClientImpl) UpdateTx(ctx context.Context, in *UpdateTxReq, opts ...http.CallOption) (*UpdateTxResp, error) {
	var out UpdateTxResp
	pattern := "/api.transaction.service.v1.Transaction/UpdateTx"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTransactionUpdateTx))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
