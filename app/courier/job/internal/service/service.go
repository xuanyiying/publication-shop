package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "github.com/publication-shop/api/courier/job/v1"
	"github.com/publication-shop/app/courier/job/internal/biz"
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
