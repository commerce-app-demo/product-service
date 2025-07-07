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

func (s *ProductServiceServer) GetProduct(ctx context.Context, req *productspb.GetProductRequest) (*productspb.Product, error) {
	product, err := s.ProductService.GetProductById(ctx, req.Id)

	if err != nil {
		return nil, err
	}

	return &productspb.Product{
		Id:    product.Id,
		Name:  product.Name,
		Price: product.Price,
	}, nil

}
