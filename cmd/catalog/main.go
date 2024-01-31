package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	"github.com/rodolfogomes/goapi/internal/database"
	"github.com/rodolfogomes/goapi/internal/service"
	"github.com/rodolfogomes/goapi/internal/webserver"
)

func main() {
	db, err := sql.Open("postgres", "postgres://root:root@localhost:5432/imersao17?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(*categoryDB)

	productDB := database.NewProductDB(db)
	productService := service.NewProductService(*productDB)

	webCategoryHandler := webserver.NewWebCategoryHandler(*categoryService)
	webProductHandler := webserver.NewWebProductHandler(*productService)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)

	c.Post("/category", webCategoryHandler.CreateCategory)
	c.Get("/category/{id}", webCategoryHandler.GetCategory)
	c.Get("/categories", webCategoryHandler.GetCategories)

	c.Post("/product", webProductHandler.CreateProduct)
	c.Get("/product/{id}", webProductHandler.GetProduct)
	c.Get("/products", webProductHandler.GetProducts)
	c.Get("/products/category/{category_id}", webProductHandler.GetProductByCategoryID)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", c)

}
