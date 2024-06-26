// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.4

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

type ShopAdminHTTPServer interface {
	CreateBook(context.Context, *CreateBookReq) (*CreateBookReply, error)
	DeleteBook(context.Context, *DeleteBookReq) (*DeleteBookReply, error)
	GetOrder(context.Context, *GetOrderReq) (*GetOrderReply, error)
	GetUser(context.Context, *GetUserReq) (*GetUserReply, error)
	ListBook(context.Context, *ListBookReq) (*ListBookReply, error)
	ListOrder(context.Context, *ListOrderReq) (*ListOrderReply, error)
	ListUser(context.Context, *ListUserReq) (*ListUserReply, error)
	Login(context.Context, *LoginReq) (*LoginReply, error)
	Logout(context.Context, *LogoutReq) (*LogoutReply, error)
	UpdateBook(context.Context, *UpdateBookReq) (*UpdateBookReply, error)
}

func RegisterShopAdminHTTPServer(s *http.Server, srv ShopAdminHTTPServer) {
	r := s.Route("/")
	r.POST("/admin/v1/login", _ShopAdmin_Login0_HTTP_Handler(srv))
	r.POST("/admin/v1/logout", _ShopAdmin_Logout0_HTTP_Handler(srv))
	r.GET("/admin/v1/catalog/Books", _ShopAdmin_ListBook0_HTTP_Handler(srv))
	r.POST("/admin/v1/catalog/Books", _ShopAdmin_CreateBook0_HTTP_Handler(srv))
	r.PUT("/admin/v1/catalog/Books/{id}", _ShopAdmin_UpdateBook0_HTTP_Handler(srv))
	r.DELETE("/admin/v1/catalog/Books/{id}", _ShopAdmin_DeleteBook0_HTTP_Handler(srv))
	r.GET("/admin/v1/orders", _ShopAdmin_ListOrder0_HTTP_Handler(srv))
	r.GET("/admin/v1/orders", _ShopAdmin_GetOrder0_HTTP_Handler(srv))
	r.GET("/admin/v1/users", _ShopAdmin_ListUser0_HTTP_Handler(srv))
	r.GET("/admin/v1/users/{id}", _ShopAdmin_GetUser0_HTTP_Handler(srv))
}

func _ShopAdmin_Login0_HTTP_Handler(srv ShopAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.admin.v1.ShopAdmin/Login")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _ShopAdmin_Logout0_HTTP_Handler(srv ShopAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LogoutReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.admin.v1.ShopAdmin/Logout")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Logout(ctx, req.(*LogoutReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LogoutReply)
		return ctx.Result(200, reply)
	}
}

func _ShopAdmin_ListBook0_HTTP_Handler(srv ShopAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListBookReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.admin.v1.ShopAdmin/ListBook")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListBook(ctx, req.(*ListBookReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListBookReply)
		return ctx.Result(200, reply)
	}
}

func _ShopAdmin_CreateBook0_HTTP_Handler(srv ShopAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateBookReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.admin.v1.ShopAdmin/CreateBook")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateBook(ctx, req.(*CreateBookReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateBookReply)
		return ctx.Result(200, reply)
	}
}

func _ShopAdmin_UpdateBook0_HTTP_Handler(srv ShopAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateBookReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.admin.v1.ShopAdmin/UpdateBook")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateBook(ctx, req.(*UpdateBookReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateBookReply)
		return ctx.Result(200, reply)
	}
}

func _ShopAdmin_DeleteBook0_HTTP_Handler(srv ShopAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteBookReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.admin.v1.ShopAdmin/DeleteBook")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteBook(ctx, req.(*DeleteBookReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteBookReply)
		return ctx.Result(200, reply)
	}
}

func _ShopAdmin_ListOrder0_HTTP_Handler(srv ShopAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListOrderReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.admin.v1.ShopAdmin/ListOrder")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListOrder(ctx, req.(*ListOrderReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListOrderReply)
		return ctx.Result(200, reply)
	}
}

func _ShopAdmin_GetOrder0_HTTP_Handler(srv ShopAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetOrderReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.admin.v1.ShopAdmin/GetOrder")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetOrder(ctx, req.(*GetOrderReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetOrderReply)
		return ctx.Result(200, reply)
	}
}

func _ShopAdmin_ListUser0_HTTP_Handler(srv ShopAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUserReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.admin.v1.ShopAdmin/ListUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUser(ctx, req.(*ListUserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserReply)
		return ctx.Result(200, reply)
	}
}

func _ShopAdmin_GetUser0_HTTP_Handler(srv ShopAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.admin.v1.ShopAdmin/GetUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetUserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserReply)
		return ctx.Result(200, reply)
	}
}

type ShopAdminHTTPClient interface {
	CreateBook(ctx context.Context, req *CreateBookReq, opts ...http.CallOption) (rsp *CreateBookReply, err error)
	DeleteBook(ctx context.Context, req *DeleteBookReq, opts ...http.CallOption) (rsp *DeleteBookReply, err error)
	GetOrder(ctx context.Context, req *GetOrderReq, opts ...http.CallOption) (rsp *GetOrderReply, err error)
	GetUser(ctx context.Context, req *GetUserReq, opts ...http.CallOption) (rsp *GetUserReply, err error)
	ListBook(ctx context.Context, req *ListBookReq, opts ...http.CallOption) (rsp *ListBookReply, err error)
	ListOrder(ctx context.Context, req *ListOrderReq, opts ...http.CallOption) (rsp *ListOrderReply, err error)
	ListUser(ctx context.Context, req *ListUserReq, opts ...http.CallOption) (rsp *ListUserReply, err error)
	Login(ctx context.Context, req *LoginReq, opts ...http.CallOption) (rsp *LoginReply, err error)
	Logout(ctx context.Context, req *LogoutReq, opts ...http.CallOption) (rsp *LogoutReply, err error)
	UpdateBook(ctx context.Context, req *UpdateBookReq, opts ...http.CallOption) (rsp *UpdateBookReply, err error)
}

type ShopAdminHTTPClientImpl struct {
	cc *http.Client
}

func NewShopAdminHTTPClient(client *http.Client) ShopAdminHTTPClient {
	return &ShopAdminHTTPClientImpl{client}
}

func (c *ShopAdminHTTPClientImpl) CreateBook(ctx context.Context, in *CreateBookReq, opts ...http.CallOption) (*CreateBookReply, error) {
	var out CreateBookReply
	pattern := "/admin/v1/catalog/Books"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/shop.admin.v1.ShopAdmin/CreateBook"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopAdminHTTPClientImpl) DeleteBook(ctx context.Context, in *DeleteBookReq, opts ...http.CallOption) (*DeleteBookReply, error) {
	var out DeleteBookReply
	pattern := "/admin/v1/catalog/Books/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/shop.admin.v1.ShopAdmin/DeleteBook"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopAdminHTTPClientImpl) GetOrder(ctx context.Context, in *GetOrderReq, opts ...http.CallOption) (*GetOrderReply, error) {
	var out GetOrderReply
	pattern := "/admin/v1/orders"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/shop.admin.v1.ShopAdmin/GetOrder"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopAdminHTTPClientImpl) GetUser(ctx context.Context, in *GetUserReq, opts ...http.CallOption) (*GetUserReply, error) {
	var out GetUserReply
	pattern := "/admin/v1/users/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/shop.admin.v1.ShopAdmin/GetUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopAdminHTTPClientImpl) ListBook(ctx context.Context, in *ListBookReq, opts ...http.CallOption) (*ListBookReply, error) {
	var out ListBookReply
	pattern := "/admin/v1/catalog/Books"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/shop.admin.v1.ShopAdmin/ListBook"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopAdminHTTPClientImpl) ListOrder(ctx context.Context, in *ListOrderReq, opts ...http.CallOption) (*ListOrderReply, error) {
	var out ListOrderReply
	pattern := "/admin/v1/orders"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/shop.admin.v1.ShopAdmin/ListOrder"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopAdminHTTPClientImpl) ListUser(ctx context.Context, in *ListUserReq, opts ...http.CallOption) (*ListUserReply, error) {
	var out ListUserReply
	pattern := "/admin/v1/users"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/shop.admin.v1.ShopAdmin/ListUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopAdminHTTPClientImpl) Login(ctx context.Context, in *LoginReq, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/admin/v1/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/shop.admin.v1.ShopAdmin/Login"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopAdminHTTPClientImpl) Logout(ctx context.Context, in *LogoutReq, opts ...http.CallOption) (*LogoutReply, error) {
	var out LogoutReply
	pattern := "/admin/v1/logout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/shop.admin.v1.ShopAdmin/Logout"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopAdminHTTPClientImpl) UpdateBook(ctx context.Context, in *UpdateBookReq, opts ...http.CallOption) (*UpdateBookReply, error) {
	var out UpdateBookReply
	pattern := "/admin/v1/catalog/Books/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/shop.admin.v1.ShopAdmin/UpdateBook"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
