package service

import (
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/go-kratos/publication-shop/api/courier/job/v1"
	"github.com/go-kratos/publication-shop/app/courier/job/internal/biz"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewCourierService)

type CourierService struct {
	v1.UnimplementedCourierServer

	oc  *biz.CourierUseCase
	log *log.Helper
}

func NewCourierService(oc *biz.CourierUseCase, logger log.Logger) *CourierService {
	return &CourierService{
		oc:  oc,
		log: log.NewHelper(log.With(logger, "module", "service/courier"))}
}
