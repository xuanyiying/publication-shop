package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Book struct {
	Id          int64
	title       string
	author      string
	price       float64
	stock       int64
	detail      string
	isbn        string
	publishDate string
}

type BookRepo interface {
	Create(ctx context.Context, c *Book) (*Book, error)
	Get(ctx context.Context, id int64) (*Book, error)
	Update(ctx context.Context, c *Book) (*Book, error)
	List(ctx context.Context, pageNum, pageSize int64) ([]*Book, error)
	ListByIsbn(ctx context.Context, isbn string) ([]*Book, error)
	ListByKeyword(ctx context.Context, keyword string) ([]*Book, error)
}

type BookUseCase struct {
	repo BookRepo
	log  *log.Helper
}

func NewBookUseCase(repo BookRepo, logger log.Logger) *BookUseCase {
	return &BookUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/book"))}
}

func (uc *BookUseCase) Create(ctx context.Context, u *Book) (*Book, error) {
	return uc.repo.Create(ctx, u)
}

func (uc *BookUseCase) Get(ctx context.Context, id int64) (*Book, error) {
	return uc.repo.Get(ctx, id)
}

func (uc *BookUseCase) Update(ctx context.Context, u *Book) (*Book, error) {
	return uc.repo.Update(ctx, u)
}

func (uc *BookUseCase) List(ctx context.Context, pageNum, pageSize int64) ([]*Book, error) {
	return uc.repo.List(ctx, pageNum, pageSize)
}
func (uc *BookUseCase) ListByIsbn(ctx context.Context, isbn string) ([]*Book, error) {
	return uc.repo.ListByIsbn(ctx, isbn)
}

func (uc *BookUseCase) ListByKeyword(ctx context.Context, keyword string) ([]*Book, error) {
	return uc.repo.ListByKeyword(ctx, keyword)
}
