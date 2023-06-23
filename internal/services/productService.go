package services

import (
	"productApi/internal/dtos"
	"productApi/internal/entity"
	"productApi/internal/repository"
	"strings"
)

type ProductService struct {

	// Here we will pass productRepo as dependency
	ProductRepository *repository.ProductRepository
}

func (pr *ProductService) DeleteProductById(productId string) error {

	err := pr.ProductRepository.DeleteProduct(productId)
	if err != nil {
		return err
	}

	return nil

}

func (pr *ProductService) GetProductById(productId string) (*dtos.ProductMapperEntity, error) {

	product, err := pr.ProductRepository.GetProduct(productId)
	if err != nil {
		return nil, err
	}

	// We will Convert the entity to Dto

	pm := dtos.ProductMapperEntity{
		Id:       product.Id,
		Title:    product.Title,
		Year:     product.Year,
		Director: product.Director,
		Cast:     strings.Join(product.Cast, ","),
	}

	return &pm, nil

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
