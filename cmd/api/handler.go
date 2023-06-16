package api

import (
	"fmt"
	"net/http"
	"productApi/internal/services"
)

type Handler struct {
	// we will give the dependency of service
	ProductService *services.ProductService // dependency
}

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	h.ProductService.GetAllProducts()

	fmt.Fprintf(w, "Hello world")

}
