package data

import (
	"context"

	cartv1 "github.com/xuanyiying/publication-shop/api/cart/service/v1"
	catalogv1 "github.com/xuanyiying/publication-shop/api/catalog/service/v1"
	orderv1 "github.com/xuanyiying/publication-shop/api/order/service/v1"
	paymentv1 "github.com/xuanyiying/publication-shop/api/payment/service/v1"
	userv1 "github.com/xuanyiying/publication-shop/api/user/service/v1"
	"github.com/xuanyiying/publication-shop/app/shop/interface/internal/conf"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwt2 "github.com/golang-jwt/jwt/v5"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewDiscovery,
	NewRegistrar,
	NewUserServiceClient,
	NewCartServiceClient,
	NewCatalogServiceClient,
	NewOrderServiceClient,
	NewPaymentServiceClient,
	NewUserRepo,
	NewBookRepo,
)

// Data .
type Data struct {
	log *log.Helper
	uc  userv1.UserClient
	cc  cartv1.CartClient
	bc  catalogv1.CatalogClient
}

// NewData .
func NewData(
	conf *conf.Data,
	logger log.Logger,
	uc userv1.UserClient,
	cc cartv1.CartClient,
	bc catalogv1.CatalogClient,
) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &Data{log: l, uc: uc, cc: cc, bc: bc}, nil
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewUserServiceClient(ac *conf.Auth, r registry.Discovery, tp *tracesdk.TracerProvider) userv1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///Book.user.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
			jwt.Client(func(token *jwt2.Token) (interface{}, error) {
				return []byte(ac.ServiceKey), nil
			}, jwt.WithSigningMethod(jwt2.SigningMethodHS256)),
		),
	)
	if err != nil {
		panic(err)
	}
	c := userv1.NewUserClient(conn)
	return c
}

func NewCartServiceClient(r registry.Discovery, tp *tracesdk.TracerProvider) cartv1.CartClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///Book.cart.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return cartv1.NewCartClient(conn)
}

func NewCatalogServiceClient(r registry.Discovery, tp *tracesdk.TracerProvider) catalogv1.CatalogClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///Book.catalog.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return catalogv1.NewCatalogClient(conn)
}

func NewOrderServiceClient(r registry.Discovery, tp *tracesdk.TracerProvider) orderv1.OrderClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///Book.order.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return orderv1.NewOrderClient(conn)
}

func NewPaymentServiceClient(r registry.Discovery, tp *tracesdk.TracerProvider) paymentv1.PaymentClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///Book.payment.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return paymentv1.NewPaymentClient(conn)
}
