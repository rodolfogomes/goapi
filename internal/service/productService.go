package service

import (
	"github.com/rodolfogomes/goapi/internal/database"
	"github.com/rodolfogomes/goapi/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(productDB database.ProductDB) *ProductService {
	return &ProductService{ProductDB: productDB}
}

func (p *ProductService) CreateProduct(name, description string, price float64, categoryID, imageURL string) (*entity.Product, error) {
	product := entity.NewProduct(name, description, price, categoryID, imageURL)
	_, err := p.ProductDB.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := p.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductService) GetProduct(id string) (*entity.Product, error) {
	product, err := p.ProductDB.GetProduct(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductService) GetProductByCategory(categoryID string) ([]*entity.Product, error) {
	products, err := p.ProductDB.GetProductByCategory(categoryID)
	if err != nil {
		return nil, err
	}
	return products, nil
}
