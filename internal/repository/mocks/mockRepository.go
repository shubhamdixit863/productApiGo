package mocks

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"productApi/internal/entity"
)

type MockRepository struct {
	// Db connection would
	Db *sqlx.DB
}

func (pr *MockRepository) UpdateProduct(productId string, product entity.Product) error {

	return nil
}

func (pr *MockRepository) DeleteProduct(productId string) error {

	return nil
}

// For getting the document by ID

func (pr *MockRepository) GetProduct(productId string) (entity.Product, error) {

	return entity.Product{}, nil
}

func (pr *MockRepository) SaveProduct(product entity.Product) error {

	return nil

}

func (pr *MockRepository) GetAll() ([]entity.Product, error) {

	fmt.Println("I am called")
	var product []entity.Product

	product = append(product, entity.Product{
		Id:       "777",
		Title:    "Random",
		Year:     199,
		Director: "Stephen grider",
		Cast:     nil,
	})

	return product, nil
}
