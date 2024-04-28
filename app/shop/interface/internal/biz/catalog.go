package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Image struct {
	URL string
}

type Book struct {
	Id          int64
	Name        string
	Description string
	Count       int64
	Images      []Image
}

type CatalogRepo interface {
	GetBook(ctx context.Context, id int64) (*Book, error)
	ListBook(ctx context.Context, pageNum, pageSize int64) ([]*Book, error)
}

type CatalogUseCase struct {
	repo CatalogRepo
	log  *log.Helper
}

func NewCatalogUseCase(repo CatalogRepo, logger log.Logger) *CatalogUseCase {
	return &CatalogUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/Book"))}
}

func (uc *CatalogUseCase) GetBook(ctx context.Context, id int64) (*Book, error) {
	return uc.repo.GetBook(ctx, id)
}

func (uc *CatalogUseCase) ListBook(ctx context.Context, pageNum, pageSize int64) ([]*Book, error) {
	return uc.repo.ListBook(ctx, pageNum, pageSize)
}
