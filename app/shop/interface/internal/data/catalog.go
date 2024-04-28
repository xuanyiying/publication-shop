package data

import (
	"context"
	"fmt"
	"golang.org/x/sync/singleflight"

	ctV1 "github.com/xuanyiying/publication-shop/api/book/service/v1"
	"github.com/xuanyiying/publication-shop/app/shop/interface/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.CatalogRepo = (*catalogRepo)(nil)

type catalogRepo struct {
	data *Data
	log  *log.Helper
	sg   *singleflight.Group
}

func NewBookRepo(data *Data, logger log.Logger) biz.CatalogRepo {
	return &catalogRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/Book")),
		sg:   &singleflight.Group{},
	}
}

func (r *catalogRepo) GetBook(ctx context.Context, id int64) (*biz.Book, error) {
	result, err, _ := r.sg.Do(fmt.Sprintf("get_Book_by_id_%d", id), func() (interface{}, error) {
		reply, err := r.data.bc.GetBook(ctx, &ctV1.GetBookReq{
			Id: id,
		})
		if err != nil {
			return nil, err
		}
		images := make([]biz.Image, 0)
		for _, x := range reply.Image {
			images = append(images, biz.Image{URL: x.Url})
		}
		return &biz.Book{
			Id:          reply.Id,
			Name:        reply.Name,
			Description: reply.Description,
			Count:       reply.Count,
			Images:      images,
		}, err
	})
	if err != nil {
		return nil, err
	}
	return result.(*biz.Book), nil
}

func (r *catalogRepo) ListBook(ctx context.Context, pageNum, pageSize int64) ([]*biz.Book, error) {
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
