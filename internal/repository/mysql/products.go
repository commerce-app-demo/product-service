package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/commerce-app-demo/product-service/internal/config"
	"github.com/commerce-app-demo/product-service/internal/models/products"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(ctx context.Context, cfg *config.DatabaseConfig) (*ProductRepository, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err := sql.Open(cfg.Driver, dsn)

	if err != nil {
		log.Fatalf("Error when opening database: %s", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &ProductRepository{
		db: db,
	}, nil
}

func (r *ProductRepository) Products() ([]products.ProductEntity, error) {
	productList := []products.ProductEntity{{
		Id:    "0",
		Name:  "Bottle",
		Price: 3000,
	}}

	return productList, nil
}

func (r *ProductRepository) ProductById(id string) (*products.ProductEntity, error) {
	product := products.ProductEntity{
		Id:    id,
		Name:  "Bottle",
		Price: 3000,
	}

	return &product, nil
}
