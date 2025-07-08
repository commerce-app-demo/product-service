package server

import (
	"context"
	"fmt"

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

func (s *ProductServiceServer) CreateProduct(ctx context.Context, req *productspb.CreateProductRequest) (*productspb.Product, error) {
	if !isRequestValid(req) {
		return nil, fmt.Errorf("%s", "Invalid request")
	}

	product, err := s.ProductService.CreateProduct(req)

	if err != nil {
		return nil, err
	}

	return product, err
}

func isRequestValid(req *productspb.CreateProductRequest) bool {
	if req.Name == "" || req.Price == 0 {
		return false
	}
	return true
}
