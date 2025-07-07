package service

import (
	"github.com/commerce-app-demo/product-service/internal/repository/mysql"
	productspb "github.com/commerce-app-demo/product-service/proto"
)

type ProductService struct {
	Repo *mysql.ProductRepository
}

func (s *ProductService) GetProducts() (*productspb.ProductArray, error) {
	productArray, err := s.Repo.Products()
	if err != nil {
		return nil, err
	}

	var productArrayPb []*productspb.Product

	for _, p := range productArray {
		productPb := &productspb.Product{
			Id:    p.Id,
			Name:  p.Name,
			Price: p.Price,
		}

		productArrayPb = append(productArrayPb, productPb)
	}

	return &productspb.ProductArray{
		Products: productArrayPb,
	}, nil
}

func (s *ProductService) GetProductById(id string) (*productspb.Product, error) {
	// through golang magic, calling a function when repo is nil is totally okay
	// but if i were to do something like the commented code below, it errors because i try to dereference a nil object...
	// log.Printf("Number, totally unrelated: %d\n", s.repo.Number)

	product, err := s.Repo.ProductById(id)

	if err != nil {
		return nil, err
	}

	return &productspb.Product{
		Id:    product.Id,
		Name:  product.Name,
		Price: product.Price,
	}, nil
}
