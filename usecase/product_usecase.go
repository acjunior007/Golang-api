package usecase

import (
	"errors"
	"go-api/model"
	"go-api/repository"
)

// ErrInvalidProductData is returned when product data is invalid.
var ErrInvalidProductData = errors.New("invalid product data")

// ... your existing

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {

	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	if product.Name == "" || product.Price <= 0 {
		return model.Product{}, repository.ErrInvalidProductData
	}
	id, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}
	product.ID = id
	return product, nil
}

func (pu *ProductUsecase) GetProductByID(id int) (*model.Product, error) {
	product, err := pu.repository.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
