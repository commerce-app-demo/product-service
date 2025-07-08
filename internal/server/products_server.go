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
	err := validateProductCreation(req)

	if err != nil {
		return nil, err
	}

	product, err := s.ProductService.CreateProduct(req)

	if err != nil {
		return nil, err
	}

	return product, err
}

func (s *ProductServiceServer) DeleteProduct(ctx context.Context, req *productspb.DeleteProductRequest) (*productspb.DeleteProductResponse, error) {
	if req.Id == "" {
		return nil, fmt.Errorf("%s", "Invalid request")
	}

	deletedProduct, err := s.ProductService.DeleteProduct(req.Id)
	if err != nil {
		return nil, err
	}

	return &productspb.DeleteProductResponse{
		Success:        true,
		DeletedProduct: deletedProduct,
	}, nil
}

func (s *ProductServiceServer) UpdateProduct(ctx context.Context, req *productspb.UpdateProductRequest) (*productspb.UpdateProductResponse, error) {

	err := validateRequestId(req.Id)

	if err != nil {
		return nil, err
	}

	updatedProduct, err := s.ProductService.UpdateProduct(req.Id, req.Product)

	if err != nil {
		return nil, err
	}

	return &productspb.UpdateProductResponse{
		Success:        true,
		UpdatedProduct: updatedProduct,
	}, nil
}

func validateProductCreation(req *productspb.CreateProductRequest) error {
	err := fmt.Errorf("%s", "Invalid request")
	if req.Name == "" {
		return err
	}

	if req.Price == 0 {
		return fmt.Errorf("%s cannot set price to 0", err)
	}

	return nil
}

func validateRequestId(id string) error {
	if id == "" {
		return fmt.Errorf("%s", "Invalid request ID")
	}

	return nil
}
