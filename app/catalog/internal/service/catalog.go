package service

import (
	"context"
	"github.com/xuanyiying/publication-shop/api/catalog/service/v1"
)

type CatalogService struct {
	v1.UnimplementedCatalogServer
}

func NewCatalogService() *CatalogService {
	return &CatalogService{}
}

func (s *CatalogService) CreateCatalog(ctx context.Context, req *v1.CreateCatalogRequest) (*v1.CreateCatalogReply, error) {
	return &v1.CreateCatalogReply{}, nil
}
func (s *CatalogService) UpdateCatalog(ctx context.Context, req *v1.UpdateCatalogRequest) (*v1.UpdateCatalogReply, error) {
	return &v1.UpdateCatalogReply{}, nil
}
func (s *CatalogService) DeleteCatalog(ctx context.Context, req *v1.DeleteCatalogRequest) (*v1.DeleteCatalogReply, error) {
	return &v1.DeleteCatalogReply{}, nil
}
func (s *CatalogService) GetCatalog(ctx context.Context, req *v1.GetCatalogRequest) (*v1.GetCatalogReply, error) {
	return &v1.GetCatalogReply{}, nil
}
func (s *CatalogService) ListCatalog(ctx context.Context, req *v1.ListCatalogRequest) (*v1.ListCatalogReply, error) {
	return &v1.ListCatalogReply{}, nil
}
