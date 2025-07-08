package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/commerce-app-demo/product-service/internal/models/products"
	productspb "github.com/commerce-app-demo/product-service/proto"
)

type ProductRepository struct {
	DB *sql.DB
}

func (r *ProductRepository) Products() ([]products.ProductEntity, error) {
	table := "products"

	query := fmt.Sprintf("SELECT id, name, price FROM %s LIMIT 50", table)
	stmt, err := r.DB.Prepare(query)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()

	var productArray []products.ProductEntity
	var product products.ProductEntity

	for rows.Next() {
		rows.Scan(&product.Id, &product.Name, &product.Price)
		productArray = append(productArray, product)
	}

	if len(productArray) < 1 {
		return nil, fmt.Errorf("%s", "Product is empty")
	}

	return productArray, nil
}

func (r *ProductRepository) ProductById(id string) (*products.ProductEntity, error) {
	table := "products"

	query := fmt.Sprintf("SELECT id, name, price FROM %s WHERE id = ?", table)
	stmt, err := r.DB.Prepare(query)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	var product products.ProductEntity

	// Guaranteed 1 result anyway
	if rows.Next() == false {
		return nil, fmt.Errorf("%s", "Item not found")
	} else {
		err = rows.Scan(&product.Id, &product.Name, &product.Price)
	}

	return &product, nil
}

func (r *ProductRepository) CreateProduct(req *productspb.CreateProductRequest) (*products.ProductEntity, error) {
	table := "products"

	execute := fmt.Sprintf("INSERT INTO %s (name,price) VALUES (?, ?)", table)
	stmt, err := r.DB.Prepare(execute)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(req.Name, req.Price)

	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &products.ProductEntity{
		Id:    fmt.Sprint(id),
		Name:  req.Name,
		Price: req.Price,
	}, nil

}
