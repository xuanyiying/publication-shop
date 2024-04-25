package service

import (
	"context"
	"github.com/publication-shop/api/shop/interface/v1"
)

func (s *ShopInterface) ListPublication(ctx context.Context, req *v1.ListPublicationReq) (*v1.ListPublicationReply, error) {
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
func (s *ShopInterface) GetPublication(ctx context.Context, req *v1.GetPublicationReq) (*v1.GetPublicationReply, error) {
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
