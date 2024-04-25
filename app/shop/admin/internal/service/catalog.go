package service

import (
	"context"
	"github.com/xuanyiying/publication-shop/api/shop/admin/v1"
)

func (s *ShopAdmin) ListPublication(ctx context.Context, req *v1.ListPublicationReq) (*v1.ListPublicationReply, error) {
	rv, err := s.cc.ListPublication(ctx, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	reply := &v1.ListPublicationReply{
		Results: make([]*v1.ListPublicationReply_Publication, 0),
	}
	for _, x := range rv {
		item := &v1.ListPublicationReply_Publication{
			Id:          x.Id,
			Name:        x.Name,
			Description: x.Description,
			Count:       x.Count,
			Image:       make([]*v1.ListPublicationReply_Publication_Image, 0),
		}
		for _, img := range x.Images {
			item.Image = append(item.Image, &v1.ListPublicationReply_Publication_Image{Url: img.URL})
		}
		reply.Results = append(reply.Results, item)
	}
	return reply, nil
}
func (s *ShopAdmin) GetPublication(ctx context.Context, req *v1.GetPublicationReq) (*v1.GetPublicationReply, error) {
	x, err := s.cc.GetPublication(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	reply := &v1.GetPublicationReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
		Count:       x.Count,
		Image:       make([]*v1.GetPublicationReply_Image, 0),
	}
	for _, img := range x.Images {
		reply.Image = append(reply.Image, &v1.GetPublicationReply_Image{Url: img.URL})
	}
	return reply, nil
}

func (s *ShopAdmin) CreatePublication(ctx context.Context, req *v1.CreatePublicationReq) (*v1.CreatePublicationReply, error) {
	return nil, nil
}

func (s *ShopAdmin) UpdatePublication(ctx context.Context, req *v1.UpdatePublicationReq) (*v1.UpdatePublicationReply, error) {
	return nil, nil
}

func (s *ShopAdmin) DeletePublication(ctx context.Context, req *v1.DeletePublicationReq) (*v1.DeletePublicationReply, error) {
	return nil, nil
}
