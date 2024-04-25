package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "github.com/publication-shop/api/cart/service/v1"
	"github.com/publication-shop/app/cart/service/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewCartService)

type CartService struct {
	v1.UnimplementedCartServer

	cc  *biz.CartUseCase
	log *log.Helper
}

func NewCartService(cc *biz.CartUseCase, logger log.Logger) *CartService {
	return &CartService{
		cc:  cc,
		log: log.NewHelper(log.With(logger, "module", "service/cart"))}
}
