package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/publication-shop/api/shop/admin/v1"
	"github.com/publication-shop/app/shop/admin/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewShopAdmin)

type ShopAdmin struct {
	v1.UnimplementedShopAdminServer

	log *log.Helper
	cc  *biz.CatalogUseCase
	uc  *biz.UserUseCase
}

func NewShopAdmin(uc *biz.UserUseCase, cc *biz.CatalogUseCase, logger log.Logger) *ShopAdmin {
	return &ShopAdmin{
		log: log.NewHelper(log.With(logger, "module", "service/interface")),
		uc:  uc,
		cc:  cc,
	}
}
