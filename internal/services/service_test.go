package services

import (
	"github.com/stretchr/testify/assert"
	"productApi/internal/repository/mocks"
	"testing"
)

func TestProductService_GetAllProducts(t *testing.T) {

	mock := mocks.MockRepository{Db: nil}

	ps := ProductService{ProductRepository: &mock}
	products, err := ps.GetAllProducts()

	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(products), 0)

}
