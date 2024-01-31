package database

import (
	"database/sql"

	"github.com/rodolfogomes/goapi/internal/entity"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (p *ProductDB) CreateProduct(product *entity.Product) (string, error) {
	_, err := p.db.Exec("INSERT INTO products (id, name, description, price, category_id, image_url) VALUES ($1, $2, $3, $4, $5, $6)", product.ID, product.Name, product.Description, product.Price, product.CategoryID, product.ImageURL)
	if err != nil {
		return "", err
	}

	return product.ID, nil
}

func (p *ProductDB) GetProduct(id string) (*entity.Product, error) {
	row := p.db.QueryRow("SELECT * FROM products WHERE id = $1", id)

	var product entity.Product
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductDB) GetProducts() ([]*entity.Product, error) {
	rows, err := p.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}

func (p *ProductDB) GetProductByCategory(categoryID string) ([]*entity.Product, error) {
	rows, err := p.db.Query("SELECT * FROM products WHERE category_id = $1", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}
