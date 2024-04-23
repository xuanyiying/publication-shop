package service

import (
	"context"

	v1 "github.com/go-kratos/publication-shop/api/catalog/service/v1"
	"github.com/go-kratos/publication-shop/app/catalog/service/internal/biz"
)

func (s *CatalogService) CreatePublication(ctx context.Context, req *v1.CreatePublicationReq) (*v1.CreatePublicationReply, error) {
	b := &biz.Publication{
		Name:        req.Name,
		Description: req.Description,
		Count:       req.Count,
		Images:      make([]biz.Image, 0),
	}
	for _, x := range req.Image {
		b.Images = append(b.Images, biz.Image{URL: x.Url})
	}
	x, err := s.bc.Create(ctx, b)
	img := make([]*v1.CreatePublicationReply_Image, 0)
	for _, i := range x.Images {
		img = append(img, &v1.CreatePublicationReply_Image{Url: i.URL})
	}
	return &v1.CreatePublicationReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
		Count:       x.Count,
		Image:       img,
	}, err
}

func (s *CatalogService) GetPublication(ctx context.Context, req *v1.GetPublicationReq) (*v1.GetPublicationReply, error) {
	x, err := s.bc.Get(ctx, req.Id)
	img := make([]*v1.GetPublicationReply_Image, 0)
	for _, i := range x.Images {
		img = append(img, &v1.GetPublicationReply_Image{Url: i.URL})
	}
	return &v1.GetPublicationReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
		Count:       x.Count,
		Image:       img,
	}, err
}

func (s *CatalogService) UpdatePublication(ctx context.Context, req *v1.UpdatePublicationReq) (*v1.UpdatePublicationReply, error) {
	b := &biz.Publication{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Count:       req.Count,
		Images:      make([]biz.Image, 0),
	}
	for _, x := range req.Image {
		b.Images = append(b.Images, biz.Image{URL: x.Url})
	}
	x, err := s.bc.Update(ctx, b)
	img := make([]*v1.UpdatePublicationReply_Image, 0)
	for _, i := range x.Images {
		img = append(img, &v1.UpdatePublicationReply_Image{Url: i.URL})
	}
	return &v1.UpdatePublicationReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
		Count:       x.Count,
		Image:       img,
	}, err
}

func (s *CatalogService) ListPublication(ctx context.Context, req *v1.ListPublicationReq) (*v1.ListPublicationReply, error) {
	rv, err := s.bc.List(ctx, req.PageNum, req.PageSize)
	rs := make([]*v1.ListPublicationReply_Publication, 0)
	for _, x := range rv {
		img := make([]*v1.ListPublicationReply_Publication_Image, 0)
		for _, i := range x.Images {
			img = append(img, &v1.ListPublicationReply_Publication_Image{Url: i.URL})
		}
		rs = append(rs, &v1.ListPublicationReply_Publication{
			Id:          x.Id,
			Name:        x.Name,
			Description: x.Description,
			Image:       img,
		})
	}
	return &v1.ListPublicationReply{
		Results: rs,
	}, err
}

func (s *CatalogService) ListPublicationNextToken(ctx context.Context, req *v1.ListPublicationNextTokenReq) (*v1.ListPublicationReplyNextToken, error) {
	rv, token, err := s.bc.ListNext(ctx, req.PageSize, req.PageToken)
	rs := make([]*v1.ListPublicationReplyNextToken_Publication, 0)
	for _, x := range rv {
		img := make([]*v1.ListPublicationReplyNextToken_Publication_Image, 0)
		for _, i := range x.Images {
			img = append(img, &v1.ListPublicationReplyNextToken_Publication_Image{Url: i.URL})
		}
		rs = append(rs, &v1.ListPublicationReplyNextToken_Publication{
			Id:          x.Id,
			Name:        x.Name,
			Description: x.Description,
			Image:       img,
		})
	}
	return &v1.ListPublicationReplyNextToken{
		Results:       rs,
		NextPageToken: token,
	}, err
}
