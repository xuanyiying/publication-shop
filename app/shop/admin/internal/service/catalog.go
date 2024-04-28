package service

import (
	"context"
	"github.com/xuanyiying/publication-shop/api/shop/admin/v1"
)

func (s *ShopAdmin) ListBook(ctx context.Context, req *v1.ListBookReq) (*v1.ListBookReply, error) {
	rv, err := s.cc.ListBook(ctx, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	reply := &v1.ListBookReply{
		Results: make([]*v1.ListBookReply_Book, 0),
	}
	for _, x := range rv {
		item := &v1.ListBookReply_Book{
			Id:          x.Id,
			Name:        x.Name,
			Description: x.Description,
			Count:       x.Count,
			Image:       make([]*v1.ListBookReply_Book_Image, 0),
		}
		for _, img := range x.Images {
			item.Image = append(item.Image, &v1.ListBookReply_Book_Image{Url: img.URL})
		}
		reply.Results = append(reply.Results, item)
	}
	return reply, nil
}
func (s *ShopAdmin) GetBook(ctx context.Context, req *v1.GetBookReq) (*v1.GetBookReply, error) {
	x, err := s.cc.GetBook(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	reply := &v1.GetBookReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
		Count:       x.Count,
		Image:       make([]*v1.GetBookReply_Image, 0),
	}
	for _, img := range x.Images {
		reply.Image = append(reply.Image, &v1.GetBookReply_Image{Url: img.URL})
	}
	return reply, nil
}

func (s *ShopAdmin) CreateBook(ctx context.Context, req *v1.CreateBookReq) (*v1.CreateBookReply, error) {
	return nil, nil
}

func (s *ShopAdmin) UpdateBook(ctx context.Context, req *v1.UpdateBookReq) (*v1.UpdateBookReply, error) {
	return nil, nil
}

func (s *ShopAdmin) DeleteBook(ctx context.Context, req *v1.DeleteBookReq) (*v1.DeleteBookReply, error) {
	return nil, nil
}
