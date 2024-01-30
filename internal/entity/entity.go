package entity

import "github.com/google/uuid"

type Category struct {
	ID   string
	Name string
}

func NewCategory(name string) *Category {
	return &Category{
		ID:   uuid.New().String(),
		Name: name,
	}
}

type Product struct {
	ID         string
	Name       string
	Desription string
	Price      float64
	CategoryID string
	ImageURL   string
}

func NewProduct(name string, description string, price float64, categoryID string, imageURL string) *Product {
	return &Product{
		ID:         uuid.New().String(),
		Name:       name,
		Desription: description,
		Price:      price,
		CategoryID: categoryID,
		ImageURL:   imageURL,
	}
}
