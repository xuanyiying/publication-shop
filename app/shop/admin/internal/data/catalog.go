package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/xuanyiying/publication-shop/app/shop/admin/internal/biz"

	catalogv1 "github.com/xuanyiying/publication-shop/api/catalog/service/v1"
)

var _ biz.CatalogRepo = (*catalogRepo)(nil)

type catalogRepo struct {
	data *Data
	log  *log.Helper
}

func NewCatalogRepo(data *Data, logger log.Logger) biz.CatalogRepo {
	return &catalogRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/Book")),
	}
}

func (r *catalogRepo) GetBook(ctx context.Context, id int64) (*biz.Book, error) {
	reply, err := r.data.bc.GetBook(ctx, &catalogv1.GetBookReq{
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
}

func (r *catalogRepo) ListBook(ctx context.Context, pageNum, pageSize int64) ([]*biz.Book, error) {
	reply, err := r.data.bc.ListBook(ctx, &catalogv1.ListBookReq{
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

func (r *catalogRepo) CreateBook(ctx context.Context, b *biz.Book) (*biz.Book, error) {
	images := make([]*catalogv1.CreateBookReq_Image, 0)
	for _, x := range b.Images {
		images = append(images, &catalogv1.CreateBookReq_Image{Url: x.URL})
	}
	reply, err := r.data.bc.CreateBook(ctx, &catalogv1.CreateBookReq{
		Name:        b.Name,
		Description: b.Description,
		Count:       b.Count,
		Image:       images,
	})
	if err != nil {
		return nil, err
	}

	return &biz.Book{
		Id: reply.Id,
	}, err
}

func (r *catalogRepo) UpdateBook(ctx context.Context, b *biz.Book) (*biz.Book, error) {
	images := make([]*catalogv1.UpdateBookReq_Image, 0)
	for _, x := range b.Images {
		images = append(images, &catalogv1.UpdateBookReq_Image{Url: x.URL})
	}
	reply, err := r.data.bc.UpdateBook(ctx, &catalogv1.UpdateBookReq{
		Name:        b.Name,
		Description: b.Description,
		Count:       b.Count,
		Image:       images,
	})
	if err != nil {
		return nil, err
	}

	return &biz.Book{
		Id: reply.Id,
	}, err
}
