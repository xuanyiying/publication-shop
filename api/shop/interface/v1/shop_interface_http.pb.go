// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.2

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

type ShopInterfaceHTTPServer interface {
	AddCartItem(context.Context, *AddCartItemReq) (*AddCartItemReply, error)
	CreateAddress(context.Context, *CreateAddressReq) (*CreateAddressReply, error)
	CreateCard(context.Context, *CreateCardReq) (*CreateCardReply, error)
	CreateOrder(context.Context, *CreateOrderReq) (*CreateOrderReply, error)
	DeleteCard(context.Context, *DeleteCardReq) (*DeleteCardReply, error)
	GetAddress(context.Context, *GetAddressReq) (*GetAddressReply, error)
	GetBook(context.Context, *GetBookReq) (*GetBookReply, error)
	GetCard(context.Context, *GetCardReq) (*GetCardReply, error)
	ListAddress(context.Context, *ListAddressReq) (*ListAddressReply, error)
	ListBook(context.Context, *ListBookReq) (*ListBookReply, error)
	ListCard(context.Context, *ListCardReq) (*ListCardReply, error)
	ListCartItem(context.Context, *ListCartItemReq) (*ListCartItemReply, error)
	ListOrder(context.Context, *ListOrderReq) (*ListOrderReply, error)
	Login(context.Context, *LoginReq) (*LoginReply, error)
	Logout(context.Context, *LogoutReq) (*LogoutReply, error)
	Register(context.Context, *RegisterReq) (*RegisterReply, error)
}

func RegisterShopInterfaceHTTPServer(s *http.Server, srv ShopInterfaceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/register", _ShopInterface_Register0_HTTP_Handler(srv))
	r.POST("/v1/login", _ShopInterface_Login0_HTTP_Handler(srv))
	r.POST("/v1/logout", _ShopInterface_Logout0_HTTP_Handler(srv))
	r.GET("/v1/user/addresses", _ShopInterface_ListAddress0_HTTP_Handler(srv))
	r.POST("/v1/user/addresses", _ShopInterface_CreateAddress0_HTTP_Handler(srv))
	r.GET("/v1/user/addresses/{id}", _ShopInterface_GetAddress0_HTTP_Handler(srv))
	r.GET("/v1/user/cards", _ShopInterface_ListCard0_HTTP_Handler(srv))
	r.POST("/v1/user/cards", _ShopInterface_CreateCard0_HTTP_Handler(srv))
	r.GET("/v1/user/cards/{id}", _ShopInterface_GetCard0_HTTP_Handler(srv))
	r.DELETE("/v1/user/cards/{id}", _ShopInterface_DeleteCard0_HTTP_Handler(srv))
	r.GET("/v1/catalog/Books", _ShopInterface_ListBook0_HTTP_Handler(srv))
	r.GET("/v1/catalog/Books/{id}", _ShopInterface_GetBook0_HTTP_Handler(srv))
	r.GET("/v1/cart", _ShopInterface_ListCartItem0_HTTP_Handler(srv))
	r.POST("/v1/cart", _ShopInterface_AddCartItem0_HTTP_Handler(srv))
	r.POST("/v1/orders", _ShopInterface_CreateOrder0_HTTP_Handler(srv))
	r.GET("/v1/orders", _ShopInterface_ListOrder0_HTTP_Handler(srv))
}

func _ShopInterface_Register0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.interface.v1.ShopInterface/Register")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterReply)
		return ctx.Result(200, reply)
	}
}

func _ShopInterface_Login0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.interface.v1.ShopInterface/Login")
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

func _ShopInterface_Logout0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LogoutReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.interface.v1.ShopInterface/Logout")
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

func _ShopInterface_ListAddress0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListAddressReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.interface.v1.ShopInterface/ListAddress")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListAddress(ctx, req.(*ListAddressReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListAddressReply)
		return ctx.Result(200, reply)
	}
}

func _ShopInterface_CreateAddress0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateAddressReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.interface.v1.ShopInterface/CreateAddress")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateAddress(ctx, req.(*CreateAddressReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateAddressReply)
		return ctx.Result(200, reply)
	}
}

func _ShopInterface_GetAddress0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetAddressReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.interface.v1.ShopInterface/GetAddress")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAddress(ctx, req.(*GetAddressReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetAddressReply)
		return ctx.Result(200, reply)
	}
}

func _ShopInterface_ListCard0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListCardReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.interface.v1.ShopInterface/ListCard")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListCard(ctx, req.(*ListCardReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListCardReply)
		return ctx.Result(200, reply)
	}
}

func _ShopInterface_CreateCard0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateCardReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.interface.v1.ShopInterface/CreateCard")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateCard(ctx, req.(*CreateCardReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateCardReply)
		return ctx.Result(200, reply)
	}
}

func _ShopInterface_GetCard0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCardReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.interface.v1.ShopInterface/GetCard")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCard(ctx, req.(*GetCardReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCardReply)
		return ctx.Result(200, reply)
	}
}

func _ShopInterface_DeleteCard0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteCardReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.interface.v1.ShopInterface/DeleteCard")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteCard(ctx, req.(*DeleteCardReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteCardReply)
		return ctx.Result(200, reply)
	}
}

func _ShopInterface_ListBook0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListBookReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.interface.v1.ShopInterface/ListBook")
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

func _ShopInterface_GetBook0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetBookReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.interface.v1.ShopInterface/GetBook")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetBook(ctx, req.(*GetBookReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetBookReply)
		return ctx.Result(200, reply)
	}
}

func _ShopInterface_ListCartItem0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListCartItemReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.interface.v1.ShopInterface/ListCartItem")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListCartItem(ctx, req.(*ListCartItemReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListCartItemReply)
		return ctx.Result(200, reply)
	}
}

func _ShopInterface_AddCartItem0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddCartItemReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.interface.v1.ShopInterface/AddCartItem")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddCartItem(ctx, req.(*AddCartItemReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddCartItemReply)
		return ctx.Result(200, reply)
	}
}

func _ShopInterface_CreateOrder0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateOrderReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.interface.v1.ShopInterface/CreateOrder")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateOrder(ctx, req.(*CreateOrderReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateOrderReply)
		return ctx.Result(200, reply)
	}
}

func _ShopInterface_ListOrder0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListOrderReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.interface.v1.ShopInterface/ListOrder")
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

type ShopInterfaceHTTPClient interface {
	AddCartItem(ctx context.Context, req *AddCartItemReq, opts ...http.CallOption) (rsp *AddCartItemReply, err error)
	CreateAddress(ctx context.Context, req *CreateAddressReq, opts ...http.CallOption) (rsp *CreateAddressReply, err error)
	CreateCard(ctx context.Context, req *CreateCardReq, opts ...http.CallOption) (rsp *CreateCardReply, err error)
	CreateOrder(ctx context.Context, req *CreateOrderReq, opts ...http.CallOption) (rsp *CreateOrderReply, err error)
	DeleteCard(ctx context.Context, req *DeleteCardReq, opts ...http.CallOption) (rsp *DeleteCardReply, err error)
	GetAddress(ctx context.Context, req *GetAddressReq, opts ...http.CallOption) (rsp *GetAddressReply, err error)
	GetBook(ctx context.Context, req *GetBookReq, opts ...http.CallOption) (rsp *GetBookReply, err error)
	GetCard(ctx context.Context, req *GetCardReq, opts ...http.CallOption) (rsp *GetCardReply, err error)
	ListAddress(ctx context.Context, req *ListAddressReq, opts ...http.CallOption) (rsp *ListAddressReply, err error)
	ListBook(ctx context.Context, req *ListBookReq, opts ...http.CallOption) (rsp *ListBookReply, err error)
	ListCard(ctx context.Context, req *ListCardReq, opts ...http.CallOption) (rsp *ListCardReply, err error)
	ListCartItem(ctx context.Context, req *ListCartItemReq, opts ...http.CallOption) (rsp *ListCartItemReply, err error)
	ListOrder(ctx context.Context, req *ListOrderReq, opts ...http.CallOption) (rsp *ListOrderReply, err error)
	Login(ctx context.Context, req *LoginReq, opts ...http.CallOption) (rsp *LoginReply, err error)
	Logout(ctx context.Context, req *LogoutReq, opts ...http.CallOption) (rsp *LogoutReply, err error)
	Register(ctx context.Context, req *RegisterReq, opts ...http.CallOption) (rsp *RegisterReply, err error)
}

type ShopInterfaceHTTPClientImpl struct {
	cc *http.Client
}

func NewShopInterfaceHTTPClient(client *http.Client) ShopInterfaceHTTPClient {
	return &ShopInterfaceHTTPClientImpl{client}
}

func (c *ShopInterfaceHTTPClientImpl) AddCartItem(ctx context.Context, in *AddCartItemReq, opts ...http.CallOption) (*AddCartItemReply, error) {
	var out AddCartItemReply
	pattern := "/v1/cart"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/shop.interface.v1.ShopInterface/AddCartItem"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopInterfaceHTTPClientImpl) CreateAddress(ctx context.Context, in *CreateAddressReq, opts ...http.CallOption) (*CreateAddressReply, error) {
	var out CreateAddressReply
	pattern := "/v1/user/addresses"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/shop.interface.v1.ShopInterface/CreateAddress"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopInterfaceHTTPClientImpl) CreateCard(ctx context.Context, in *CreateCardReq, opts ...http.CallOption) (*CreateCardReply, error) {
	var out CreateCardReply
	pattern := "/v1/user/cards"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/shop.interface.v1.ShopInterface/CreateCard"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopInterfaceHTTPClientImpl) CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...http.CallOption) (*CreateOrderReply, error) {
	var out CreateOrderReply
	pattern := "/v1/orders"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/shop.interface.v1.ShopInterface/CreateOrder"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopInterfaceHTTPClientImpl) DeleteCard(ctx context.Context, in *DeleteCardReq, opts ...http.CallOption) (*DeleteCardReply, error) {
	var out DeleteCardReply
	pattern := "/v1/user/cards/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/shop.interface.v1.ShopInterface/DeleteCard"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopInterfaceHTTPClientImpl) GetAddress(ctx context.Context, in *GetAddressReq, opts ...http.CallOption) (*GetAddressReply, error) {
	var out GetAddressReply
	pattern := "/v1/user/addresses/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/shop.interface.v1.ShopInterface/GetAddress"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopInterfaceHTTPClientImpl) GetBook(ctx context.Context, in *GetBookReq, opts ...http.CallOption) (*GetBookReply, error) {
	var out GetBookReply
	pattern := "/v1/catalog/Books/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/shop.interface.v1.ShopInterface/GetBook"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopInterfaceHTTPClientImpl) GetCard(ctx context.Context, in *GetCardReq, opts ...http.CallOption) (*GetCardReply, error) {
	var out GetCardReply
	pattern := "/v1/user/cards/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/shop.interface.v1.ShopInterface/GetCard"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopInterfaceHTTPClientImpl) ListAddress(ctx context.Context, in *ListAddressReq, opts ...http.CallOption) (*ListAddressReply, error) {
	var out ListAddressReply
	pattern := "/v1/user/addresses"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/shop.interface.v1.ShopInterface/ListAddress"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopInterfaceHTTPClientImpl) ListBook(ctx context.Context, in *ListBookReq, opts ...http.CallOption) (*ListBookReply, error) {
	var out ListBookReply
	pattern := "/v1/catalog/Books"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/shop.interface.v1.ShopInterface/ListBook"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopInterfaceHTTPClientImpl) ListCard(ctx context.Context, in *ListCardReq, opts ...http.CallOption) (*ListCardReply, error) {
	var out ListCardReply
	pattern := "/v1/user/cards"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/shop.interface.v1.ShopInterface/ListCard"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopInterfaceHTTPClientImpl) ListCartItem(ctx context.Context, in *ListCartItemReq, opts ...http.CallOption) (*ListCartItemReply, error) {
	var out ListCartItemReply
	pattern := "/v1/cart"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/shop.interface.v1.ShopInterface/ListCartItem"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopInterfaceHTTPClientImpl) ListOrder(ctx context.Context, in *ListOrderReq, opts ...http.CallOption) (*ListOrderReply, error) {
	var out ListOrderReply
	pattern := "/v1/orders"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/shop.interface.v1.ShopInterface/ListOrder"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopInterfaceHTTPClientImpl) Login(ctx context.Context, in *LoginReq, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/v1/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/shop.interface.v1.ShopInterface/Login"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopInterfaceHTTPClientImpl) Logout(ctx context.Context, in *LogoutReq, opts ...http.CallOption) (*LogoutReply, error) {
	var out LogoutReply
	pattern := "/v1/logout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/shop.interface.v1.ShopInterface/Logout"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopInterfaceHTTPClientImpl) Register(ctx context.Context, in *RegisterReq, opts ...http.CallOption) (*RegisterReply, error) {
	var out RegisterReply
	pattern := "/v1/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/shop.interface.v1.ShopInterface/Register"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
