package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Image struct {
	URL string
}

type Publication struct {
	Id          int64
	Name        string
	Description string
	Count       int64
	Images      []Image
}

type CatalogRepo interface {
	GetPublication(ctx context.Context, id int64) (*Publication, error)
	ListPublication(ctx context.Context, pageNum, pageSize int64) ([]*Publication, error)
}

type CatalogUseCase struct {
	repo CatalogRepo
	log  *log.Helper
}

func NewCatalogUseCase(repo CatalogRepo, logger log.Logger) *CatalogUseCase {
	return &CatalogUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/Publication"))}
}

func (uc *CatalogUseCase) GetPublication(ctx context.Context, id int64) (*Publication, error) {
	return uc.repo.GetPublication(ctx, id)
}

func (uc *CatalogUseCase) ListPublication(ctx context.Context, pageNum, pageSize int64) ([]*Publication, error) {
	return uc.repo.ListPublication(ctx, pageNum, pageSize)
}
