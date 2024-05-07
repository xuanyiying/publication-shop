package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"github.com/xuanyiying/publication-shop/app/book/internal/biz"

	v1 "github.com/xuanyiying/publication-shop/api/book/service/v1"
)

type BookService struct {
	v1.UnimplementedBookServer

	bc  *biz.BookUseCase
	log *log.Helper
}

func NewBookService(bc *biz.BookUseCase, logger log.Logger) *BookService {
	return &BookService{
		bc:  bc,
		log: log.NewHelper(log.With(logger, "module", "service/book"))}
}
func (s *BookService) ListBook(ctx context.Context, req *v1.ListBookReq) (*v1.ListBookReply, error) {
	return &v1.ListBookReply{
		Books: []*v1.BookInfo{},
	}, nil
}
func (s *BookService) CreateBook(ctx context.Context, req *v1.CreateBookReq) (*v1.CreateBookReply, error) {
	return &v1.CreateBookReply{}, nil
}
func (s *BookService) GetBook(ctx context.Context, req *v1.GetBookReq) (*v1.GetBookReply, error) {
	var books []*biz.Book
	if req.Isbn != "" {
		books, _ = s.bc.ListByIsbn(ctx, req.Isbn)
	}
	if req.BookId != 0 {
		book, _ := s.bc.Get(ctx, req.BookId)
		books = append(books, book)
	}
	if req.Keyword != "" {
		books, _ = s.bc.ListByKeyword(ctx, req.Keyword)
	}
	var bookInfoList []*v1.BookInfo
	if len(books) != 0 {
		for _, b := range books {
			var bi *v1.BookInfo
			copier.Copy(&b, bi)
			bookInfoList = append(bookInfoList, bi)
		}
	}
	return nil, nil
}
func (s *BookService) GetBookByIsbn(ctx context.Context, req *v1.GetBookReq) (*v1.GetBookReply, error) {
	return &v1.GetBookReply{}, nil
}
func (s *BookService) UpdateBookQuantity(ctx context.Context, req *v1.UpdateBookQuantityReq) (*v1.UpdateBookQuantityReply, error) {
	return &v1.UpdateBookQuantityReply{}, nil
}
