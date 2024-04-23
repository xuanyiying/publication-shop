package service

import (
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/go-kratos/publication-shop/api/catalog/service/v1"
	"github.com/go-kratos/publication-shop/app/catalog/service/internal/biz"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewCatalogService)

type CatalogService struct {
	v1.UnimplementedCatalogServer

	bc  *biz.PublicationUseCase
	log *log.Helper
}

func NewCatalogService(bc *biz.PublicationUseCase, logger log.Logger) *CatalogService {
	return &CatalogService{

		bc:  bc,
		log: log.NewHelper(log.With(logger, "module", "service/catalog"))}
}
