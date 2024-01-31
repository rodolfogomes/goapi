package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/goccy/go-json"
	"github.com/rodolfogomes/goapi/internal/entity"
	"github.com/rodolfogomes/goapi/internal/service"
)

type WebProductHandler struct {
	ProductService service.ProductService
}

func NewWebProductHandler(productService service.ProductService) *WebProductHandler {
	return &WebProductHandler{ProductService: productService}
}

func (wpH *WebProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := wpH.ProductService.CreateProduct(product.Name, product.Description, product.Price, product.CategoryID, product.ImageURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (wpH *WebProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := wpH.ProductService.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (wpH *WebProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	product, err := wpH.ProductService.GetProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func (wpH *WebProductHandler) GetProductByCategoryID(w http.ResponseWriter, r *http.Request) {

	categoryID := chi.URLParam(r, "category_id")
	if categoryID == "" {
		http.Error(w, "category_id is required", http.StatusBadRequest)
		return
	}

	products, err := wpH.ProductService.GetProductByCategory(categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}
