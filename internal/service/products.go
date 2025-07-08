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

func (s *ProductService) CreateProduct(req *productspb.CreateProductRequest) (*productspb.Product, error) {
	product, err := s.Repo.CreateProduct(req)

	if err != nil {
		return nil, err
	}

	return &productspb.Product{
		Id:    product.Id,
		Name:  product.Name,
		Price: product.Price,
	}, nil
}

func (s *ProductService) DeleteProduct(id string) (*productspb.Product, error) {
	deletedProduct, err := s.Repo.DeleteProduct(id)
	if err != nil {
		return nil, err
	}

	return &productspb.Product{
		Id:    deletedProduct.Id,
		Name:  deletedProduct.Name,
		Price: deletedProduct.Price,
	}, nil
}

func (s *ProductService) UpdateProduct(id string, req *productspb.Product) (*productspb.Product, error) {
	updatedFields := make(map[string]any)

	if req.Name != "" {
		updatedFields["name"] = req.Name
	}

	if req.Price != 0 {
		updatedFields["price"] = req.Price
	}

	updatedProduct, err := s.Repo.UpdateProduct(id, updatedFields)

	if err != nil {
		return nil, err
	}

	return &productspb.Product{
		Id:    updatedProduct.Id,
		Name:  updatedProduct.Name,
		Price: updatedProduct.Price,
	}, nil
}
