package data

import (
	"context"

	"github.com/xuanyiying/publication-shop/app/book/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.BookRepo = (*bookRepo)(nil)

type bookRepo struct {
	data *Data
	log  *log.Helper
}

func NewBookRepo(data *Data, logger log.Logger) biz.BookRepo {
	return &bookRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (r *bookRepo) Update(ctx context.Context, c *biz.Book) (*biz.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (r *bookRepo) Create(ctx context.Context, c *biz.Book) (*biz.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (r *bookRepo) Get(ctx context.Context, id int64) (*biz.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (r *bookRepo) List(ctx context.Context, pageNum, pageSize int64) ([]*biz.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (r *bookRepo) ListByIsbn(ctx context.Context, isbn string) ([]*biz.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (r *bookRepo) ListByKeyword(ctx context.Context, keyword string) ([]*biz.Book, error) {
	//TODO implement me
	panic("implement me")
}
