package service

import (
	"context"

	pb "api/order/service/v1"
)

type BookService struct {
	pb.UnimplementedBookServer
}

func NewBookService() *BookService {
	return &BookService{}
}

func (s *BookService) ListBook(ctx context.Context, req *pb.ListBookReq) (*pb.ListBookReply, error) {
	return &pb.ListBookReply{}, nil
}
func (s *BookService) CreateBook(ctx context.Context, req *pb.CreateBookReq) (*pb.CreateBookReply, error) {
	return &pb.CreateBookReply{}, nil
}
func (s *BookService) GetBook(ctx context.Context, req *pb.GetBookReq) (*pb.GetBookReply, error) {
	return &pb.GetBookReply{}, nil
}
func (s *BookService) UpdateBook(ctx context.Context, req *pb.UpdateBookReq) (*pb.UpdateBookReply, error) {
	return &pb.UpdateBookReply{}, nil
}
