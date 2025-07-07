package service

import (
	"context"

	"github.com/commerce-app-demo/product-service/internal/config"
	"github.com/commerce-app-demo/product-service/internal/models/products"
	"github.com/commerce-app-demo/product-service/internal/repository/mysql"
)

type ProductService struct {
	repo *mysql.ProductRepository
}

func (s *ProductService) GetProductById(ctx context.Context, id string) (*products.ProductEntity, error) {
	cfg := config.LoadDBConfig()

	repo, err := mysql.NewProductRepository(ctx, cfg)

	if err != nil {
		return nil, err
	}

	product, err := repo.ProductById(id)

	if err != nil {
		return nil, err
	}

	return &products.ProductEntity{
		Id:    product.Id,
		Name:  product.Name,
		Price: product.Price,
	}, nil
}
