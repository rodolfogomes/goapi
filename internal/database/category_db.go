package database

import (
	"database/sql"

	"github.com/rodolfogomes/goapi/internal/entity"
)

type CategoryDB struct {
	db *sql.DB
}

func NewCategoryDB(db *sql.DB) *CategoryDB {
	return &CategoryDB{db: db}
}

func (c *CategoryDB) CreateCategory(category *entity.Category) (string, error) {
	_, err := c.db.Exec("INSERT INTO categories (id, name) VALUES ($1, $2)", category.ID, category.Name)
	if err != nil {
		return "", err
	}

	return category.ID, nil
}

func (c *CategoryDB) GetCategory(id string) (*entity.Category, error) {
	row := c.db.QueryRow("SELECT * FROM categories WHERE id = $1", id)

	var category entity.Category
	err := row.Scan(&category.ID, &category.Name)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (c *CategoryDB) GetCategories() ([]*entity.Category, error) {
	rows, err := c.db.Query("SELECT * FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*entity.Category
	for rows.Next() {
		var category entity.Category
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	return categories, nil
}
