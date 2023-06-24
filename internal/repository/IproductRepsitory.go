package repository

import "productApi/internal/entity"

type IProductRepository interface {
	UpdateProduct(productId string, product entity.Product) error
	DeleteProduct(productId string) error
	GetProduct(productId string) (entity.Product, error)
	SaveProduct(product entity.Product) error
	GetAll() ([]entity.Product, error)
}
