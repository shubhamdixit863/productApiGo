package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"productApi/internal/dtos"
	"productApi/internal/services"
)

type Handler struct {
	// we will give the dependency of service
	ProductService *services.ProductService // dependency
}

func (h *Handler) GetProductById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		fmt.Fprintf(w, "id is missing in parameters")
		fmt.Println("id is missing in parameters")
	}

	byId, err := h.ProductService.GetProductById(id)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	marshal, err := json.Marshal(byId)
	if err != nil {
		fmt.Fprintf(w, err.Error())

	}

	// We will content type application/json
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(marshal))
}

func (h *Handler) DeleteProductById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		fmt.Fprintf(w, "id is missing in parameters")
		fmt.Println("id is missing in parameters")
	}

	err := h.ProductService.DeleteProductById(id)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	// We will content type application/json
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Product Deleted")
}

func (h *Handler) AddProduct(w http.ResponseWriter, r *http.Request) {

	// We will parse the request body
	var requestBody dtos.ProductRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Println(requestBody)
	err = h.ProductService.SaveProduct(requestBody)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, "Product Saved")
}

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.ProductService.GetAllProducts()
	if err != nil {
		fmt.Fprintf(w, err.Error())

	}

	//jsonData, err := json.Marshal(products)

	if err != nil {

		fmt.Fprintf(w, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(products)
	//fmt.Fprintf(w, string(jsonData))

}
