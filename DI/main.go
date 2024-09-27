package main

import (
	"database/sql"
	"fmt"

	"github.com/israelalvesmelo/di/product"
)

func main() {
	db, err := getDB()
	if err != nil {
		panic(err)
	}

	// Create repositories
	productRepository := product.NewProductRepository(db)

	// Create usecases
	productUsecase := product.NewProductUsecase(productRepository)

	product, err := productUsecase.GetProduct(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(product.Name)
}

func getDB() (*sql.DB, error) {
	return nil, nil
}
