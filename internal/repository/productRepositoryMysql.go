package repository

import (
	"github.com/jmoiron/sqlx"
	"productApi/internal/entity"
)

type ProductRepositoryMysql struct {
	// Db connection would
	Db *sqlx.DB
}

func (pr *ProductRepositoryMysql) UpdateProduct(productId string, product entity.Product) error {

	return nil
}

func (pr *ProductRepositoryMysql) DeleteProduct(productId string) error {

	return nil
}

// For getting the document by ID

func (pr *ProductRepositoryMysql) GetProduct(productId string) (entity.Product, error) {

	return entity.Product{}, nil
}

func (pr *ProductRepositoryMysql) SaveProduct(product entity.Product) error {

	return nil

}

func (pr *ProductRepositoryMysql) GetAll() ([]entity.Product, error) {

	return nil, nil
}
