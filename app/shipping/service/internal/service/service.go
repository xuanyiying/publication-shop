package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "github.com/publication-shop/api/shipping/service/v1"
	"github.com/publication-shop/app/shipping/service/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewShippingService)

type ShippingService struct {
	v1.UnimplementedShippingServer

	oc  *biz.ShippingUseCase
	log *log.Helper
}

func NewShippingService(oc *biz.ShippingUseCase, logger log.Logger) *ShippingService {
	return &ShippingService{
		oc:  oc,
		log: log.NewHelper(log.With(logger, "module", "service/shipping"))}
}
