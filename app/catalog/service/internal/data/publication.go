package data

import (
	"context"
	"github.com/go-kratos/publication-shop/app/catalog/service/internal/biz"
	"github.com/go-kratos/publication-shop/pkg/util/pagination"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.PublicationRepo = (*PublicationRepo)(nil)

type PublicationRepo struct {
	data *Data
	log  *log.Helper
}

func NewPublicationRepo(data *Data, logger log.Logger) biz.PublicationRepo {
	return &PublicationRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/Publication")),
	}
}

func (r *PublicationRepo) CreatePublication(ctx context.Context, b *biz.Publication) (*biz.Publication, error) {
	po, err := r.data.db.PublicationInfo.
		Create().
		SetPublicationName(b.Name).
		SetDescription(b.Description).
		SetCount(b.Count).
		SetImages(b.Images).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &biz.Publication{}, nil
}

func (r *PublicationRepo) GetPublication(ctx context.Context, id int64) (*biz.Publication, error) {
	po, err := r.data.db.PublicationInfo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.Publication{
		Id:          po.ID,
		Description: po.Description,
		Count:       po.Count,
		Images:      po.Images,
	}, nil
}

func (r *PublicationRepo) UpdatePublication(ctx context.Context, b *biz.Publication) (*biz.Publication, error) {
	po, err := r.data.db.Publication.
		Create().
		SetName(b.Name).
		SetDescription(b.Description).
		SetCount(b.Count).
		SetImages(b.Images).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &biz.Publication{
		Id:          po.ID,
		Description: po.Description,
		Count:       po.Count,
		Images:      po.Images,
	}, nil
}

func (r *PublicationRepo) ListPublication(ctx context.Context, pageNum, pageSize int64) ([]*biz.Publication, error) {
	pos, err := r.data.db.Publication.Query().
		Offset(int(pagination.GetPageOffset(pageNum, pageSize))).
		Limit(int(pageSize)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Publication, 0, len(pos))
	for _, po := range pos {
		rv = append(rv, &biz.Publication{
			Id:          po.ID,
			Description: po.Description,
			Count:       po.Count,
			Images:      po.Images,
		})
	}
	return rv, nil
}

func (r *PublicationRepo) Count(ctx context.Context) (int, error) {
	return r.data.db.Publication.Query().Count(ctx)
}

func (r *PublicationRepo) ListPublicationNext(ctx context.Context, start, end int32) ([]*biz.Publication, error) {

	pos, err := r.data.db.Publication.Query().
		Offset(int(start)).
		Limit(int(end - start)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Publication, 0, len(pos))
	for _, po := range pos {
		rv = append(rv, &biz.Publication{
			Id:          po.ID,
			Description: po.Description,
			Count:       po.Count,
			Images:      po.Images,
		})
	}
	return rv, nil
}
