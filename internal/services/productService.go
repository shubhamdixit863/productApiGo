package services

import (
	"productApi/internal/dtos"
	"productApi/internal/entity"
	"productApi/internal/repository"
)

type ProductService struct {

	// Here we will pass productRepo as dependency
	ProductRepository *repository.ProductRepository
}

func (pr *ProductService) GetAllProducts() ([]entity.Product, error) {

	// We create a dto here

	return pr.ProductRepository.GetAll()

}

func (pr *ProductService) SaveProduct(request dtos.ProductRequest) error {

	// We create a dto here

	productEntity := entity.Product{

		Title:    request.Title,
		Year:     request.Year,
		Director: request.Director,
		Cast:     request.Cast,
	}

	return pr.ProductRepository.SaveProduct(productEntity)

}
