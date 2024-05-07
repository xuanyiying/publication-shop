package data

import (
	"context"
	"fmt"
	"golang.org/x/sync/singleflight"

	ctV1 "github.com/xuanyiying/publication-shop/api/book/service/v1"
	"github.com/xuanyiying/publication-shop/app/shop/interface/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.CatalogRepo = (*bookRepo)(nil)

type bookRepo struct {
	data *Data
	log  *log.Helper
	sg   *singleflight.Group
}

func NewBookRepo(data *Data, logger log.Logger) biz.CatalogRepo {
	return &bookRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/Book")),
		sg:   &singleflight.Group{},
	}
}

func (r *bookRepo) GetBook(ctx context.Context, id int64) (*biz.Book, error) {
	result, err, _ := r.sg.Do(fmt.Sprintf("get_Book_by_id_%d", id), func() (interface{}, error) {
		reply, err := r.data.bc.GetBook(ctx, &ctV1.GetBookReq{
			Id: id,
		})
		if err != nil {
			return nil, err
		}
		images := make([]biz.Image, 0)

		return &biz.Book{
			Id:           reply.BookId,
			Title:        reply.Title,
			Description:  reply.Description,
			Quantity:     reply.Quantity,
			Images:       images,
			Price:        reply.Price,
			Author:       reply.Author,
			Category:     reply.Category,
			Publisher:    reply.Publisher,
			PageCount:    reply.PageCount,
			Isbn:         reply.Isbn,
			Language:     reply.Language,
			PublishDate:  reply.PublishDate,
			PublishTimes: reply.PublishTimes,
		}, err
	})
	if err != nil {
		return nil, err
	}
	return result.(*biz.Book), nil
}

func (r *bookRepo) ListBook(ctx context.Context, pageNum, pageSize int64) ([]*biz.Book, error) {
	reply, err := r.data.bc.ListBook(ctx, &ctV1.ListBookReq{
		PageNum:  pageNum,
		PageSize: pageSize,
	})
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Book, 0)
	for _, x := range reply.Results {
		images := make([]biz.Image, 0)
		for _, img := range x.Image {
			images = append(images, biz.Image{URL: img.Url})
		}
		rv = append(rv, &biz.Book{
			Id:          x.Id,
			Description: x.Description,
			Count:       x.Count,
			Images:      images,
		})
	}
	return rv, err
}
