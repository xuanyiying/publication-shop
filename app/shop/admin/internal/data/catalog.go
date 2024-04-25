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
		log:  log.NewHelper(log.With(logger, "module", "data/Publication")),
	}
}

func (r *catalogRepo) GetPublication(ctx context.Context, id int64) (*biz.Publication, error) {
	reply, err := r.data.bc.GetPublication(ctx, &catalogv1.GetPublicationReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	images := make([]biz.Image, 0)
	for _, x := range reply.Image {
		images = append(images, biz.Image{URL: x.Url})
	}
	return &biz.Publication{
		Id:          reply.Id,
		Name:        reply.Name,
		Description: reply.Description,
		Count:       reply.Count,
		Images:      images,
	}, err
}

func (r *catalogRepo) ListPublication(ctx context.Context, pageNum, pageSize int64) ([]*biz.Publication, error) {
	reply, err := r.data.bc.ListPublication(ctx, &catalogv1.ListPublicationReq{
		PageNum:  pageNum,
		PageSize: pageSize,
	})
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Publication, 0)
	for _, x := range reply.Results {
		images := make([]biz.Image, 0)
		for _, img := range x.Image {
			images = append(images, biz.Image{URL: img.Url})
		}
		rv = append(rv, &biz.Publication{
			Id:          x.Id,
			Description: x.Description,
			Count:       x.Count,
			Images:      images,
		})
	}
	return rv, err
}

func (r *catalogRepo) CreatePublication(ctx context.Context, b *biz.Publication) (*biz.Publication, error) {
	images := make([]*catalogv1.CreatePublicationReq_Image, 0)
	for _, x := range b.Images {
		images = append(images, &catalogv1.CreatePublicationReq_Image{Url: x.URL})
	}
	reply, err := r.data.bc.CreatePublication(ctx, &catalogv1.CreatePublicationReq{
		Name:        b.Name,
		Description: b.Description,
		Count:       b.Count,
		Image:       images,
	})
	if err != nil {
		return nil, err
	}

	return &biz.Publication{
		Id: reply.Id,
	}, err
}

func (r *catalogRepo) UpdatePublication(ctx context.Context, b *biz.Publication) (*biz.Publication, error) {
	images := make([]*catalogv1.UpdatePublicationReq_Image, 0)
	for _, x := range b.Images {
		images = append(images, &catalogv1.UpdatePublicationReq_Image{Url: x.URL})
	}
	reply, err := r.data.bc.UpdatePublication(ctx, &catalogv1.UpdatePublicationReq{
		Name:        b.Name,
		Description: b.Description,
		Count:       b.Count,
		Image:       images,
	})
	if err != nil {
		return nil, err
	}

	return &biz.Publication{
		Id: reply.Id,
	}, err
}
