package biz

import (
	"github.com/google/wire"
	_ "github.com/xuanyiying/publication-shop/api/order/service/v1"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewCatalogService)
