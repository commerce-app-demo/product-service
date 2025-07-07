package mysql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/commerce-app-demo/product-service/internal/models/products"
)

type ProductRepository struct {
	DB *sql.DB
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
	table := "products"

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", table)
	stmt, err := r.DB.Prepare(query)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(id)
	var product products.ProductEntity

	for rows.Next() {
		err = rows.Scan(&product.Id, &product.Name, &product.Price)
		if err != nil {
			log.Printf("Error when scanning: %s", err)
			return nil, err
		}
	}

	return &product, nil
}
