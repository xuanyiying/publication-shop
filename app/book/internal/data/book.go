package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"github.com/xuanyiying/publication-shop/app/book/internal/biz"
	"github.com/xuanyiying/publication-shop/app/book/internal/data/ent"
	"github.com/xuanyiying/publication-shop/app/book/internal/data/ent/book"
	"github.com/xuanyiying/publication-shop/app/book/internal/data/ent/bookdetail"
	"github.com/xuanyiying/publication-shop/pkg/util/pagination"
	"golang.org/x/sync/singleflight"
)

var _ biz.BookRepo = (*bookRepo)(nil)

type bookRepo struct {
	data *Data
	log  *log.Helper
	sg   *singleflight.Group
}

func (r *bookRepo) Get(ctx context.Context, id int64) (*biz.Book, error) {
	b, err := r.data.db.Book.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.Book{
		Id:          b.ID,
		Title:       b.Title,
		Author:      b.Author,
		Publisher:   b.Publisher,
		Isbn:        b.Isbn,
		Price:       b.Price,
		Language:    b.Language,
		Category:    b.Category,
		Description: b.Description,
		CoverImage:  b.CoverImage,
	}, nil
}

func (r *bookRepo) List(ctx context.Context, pageNum, pageSize int64) ([]*biz.Book, error) {
	pos, err := r.data.db.Book.Query().
		Offset(int(pagination.GetPageOffset(pageNum, pageSize))).
		Limit(int(pageSize)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Book, 0, len(pos))
	for _, po := range pos {
		d, _ := r.data.db.BookDetail.Query().Where(bookdetail.BookIDEQ(po.ID)).First(ctx)
		b := &biz.Book{}
		copier.Copy(b, po)
		b.Detail = d.DetailHTML
		rv = append(rv, b)
	}
	return rv, nil
}

func (r *bookRepo) Create(ctx context.Context, b *biz.Book) (*biz.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (r *bookRepo) Update(ctx context.Context, b *biz.Book) (*biz.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (r *bookRepo) ListByIsbn(ctx context.Context, isbn string) ([]*biz.Book, error) {
	books, err := r.data.db.Book.Query().
		Where(book.Isbn(isbn)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return copy(books)
}

// ListByKeyword 模糊查询
func (r *bookRepo) ListByKeyword(ctx context.Context, keyword string) ([]*biz.Book, error) {
	// 使用ent的查询构建器
	title := book.TitleContains(keyword)
	author := book.AuthorContains(keyword)
	description := book.DescriptionContains(keyword)
	publisher := book.PublisherContains(keyword)
	isbn := book.IsbnContains(keyword)
	books, err := r.data.db.Book.Query().
		Where(book.Or(title, author, publisher, description, isbn)). // 模糊查询title字段
		All(ctx)
	if err != nil {
		return nil, err
	}
	// 将ent实体转换为biz结构体
	return copy(books)
}

func copy(books []*ent.Book) ([]*biz.Book, error) {
	var bizBooks []*biz.Book
	for _, book := range books {
		value := &biz.Book{}
		copier.Copy(book, value)
		bizBooks = append(bizBooks, value) // 假设有一个转换函数
	}
	return bizBooks, nil
}
