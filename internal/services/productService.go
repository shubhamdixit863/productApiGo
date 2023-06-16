package services

import (
	"productApi/internal/repository"
)

type ProductService struct {

	// Here we will pass productRepo as dependency
	ProductRepository *repository.ProductRepository
}

func (pr *ProductService) GetAllProducts() {

	pr.ProductRepository.GetAll()

}
