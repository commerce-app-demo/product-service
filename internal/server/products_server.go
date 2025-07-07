package server

import (
	"context"

	"github.com/commerce-app-demo/product-service/internal/service"
	productspb "github.com/commerce-app-demo/product-service/proto"
)

type ProductServiceServer struct {
	productspb.UnimplementedProductServiceServer
	ProductService *service.ProductService
}

func (s *ProductServiceServer) GetProducts(ctx context.Context, req *productspb.Empty) (*productspb.ProductArray, error) {
	products, err := s.ProductService.GetProducts()

	if err != nil {
		return nil, err
	}

	return products, err
}

func (s *ProductServiceServer) GetProduct(ctx context.Context, req *productspb.GetProductRequest) (*productspb.Product, error) {
	product, err := s.ProductService.GetProductById(req.Id)

	if err != nil {
		return nil, err
	}

	return product, err

}
