package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type BookImage struct {
	ImageUrl string
	MainFlag bool
}

type Book struct {
	Id          int
	Title       string
	Author      string
	Price       float64
	Quantity    int32
	Detail      string
	Isbn        string
	PublishDate int16
	Description string
	Images      []BookImage
	Category    string
	Language    string
	PageCount   int32
	Edition     int8
	Translator  string
	CategoryId  int
	Publisher   string
	PublisherId int
	AuthorId    int
	CoverImage  string
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
