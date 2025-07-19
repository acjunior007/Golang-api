package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"go-api/model"
)

var ErrInvalidProductData = errors.New("invalid product data")
var ErrUpdateProductNotFound = errors.New("id different of the product found to update")

// ... your existing code

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return ProductRepository{
		connection: db,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {

	query := "SELECT id, name, price FROM products"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(&productObj.ID, &productObj.Name, &productObj.Price)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return []model.Product{}, err
		}
		productList = append(productList, productObj)
	}
	rows.Close()
	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int
	query, err := pr.connection.Prepare("INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id")

	if err != nil {
		fmt.Println("Error preparing query:", err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return 0, err
	}
	query.Close()

	return id, nil
}

func (pr *ProductRepository) GetProductByID(id int) (*model.Product, error) {

	query, err := pr.connection.Prepare("select id, name, price from products where id = $1")
	if err != nil {
		fmt.Println("Error preparing query:", err)
		return nil, err
	}

	var product model.Product

	err = query.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product with id %d not found", id)
		}
		fmt.Println("Error executing query:", err)
		return nil, err
	}

	query.Close()
	return &product, nil
}

func (pr *ProductRepository) Delete(id int) (int, error) {

	query, err := pr.connection.Exec("delete from products where id = $1", id)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return 0, err
	}

	count, err := query.RowsAffected()

	if err != nil {
		fmt.Println("Error getting rows affected:", err)
		return 0, err
	}

	return int(count), nil
}

func (pr *ProductRepository) Update(product model.Product) (model.Product, error) {
	query, err := pr.connection.Prepare("UPDATE products SET name = $1, price = $2 WHERE id = $3")
	if err != nil {
		fmt.Println("Error preparing query:", err)
		return model.Product{}, err
	}

	_, err = query.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return model.Product{}, err
	}
	query.Close()
	return product, nil
}
