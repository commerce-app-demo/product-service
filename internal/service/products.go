package service

import (
	"context"

	"github.com/commerce-app-demo/product-service/internal/models/products"
	"github.com/commerce-app-demo/product-service/internal/repository/mysql"
)

type ProductService struct {
	Repo *mysql.ProductRepository
}

func (s *ProductService) GetProductById(ctx context.Context, id string) (*products.ProductEntity, error) {
	// through golang magic, calling a function when repo is nil is totally okay
	// but if i were to do something like the commented code below, it errors because i try to dereference a nil object...
	// log.Printf("Number, totally unrelated: %d\n", s.repo.Number)

	product, err := s.Repo.ProductById(id)

	if err != nil {
		return nil, err
	}

	return &products.ProductEntity{
		Id:    product.Id,
		Name:  product.Name,
		Price: product.Price,
	}, nil
}
