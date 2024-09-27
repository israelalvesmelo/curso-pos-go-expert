package main

import (
	"database/sql"
	"fmt"
)

func main() {
	db, err := getDB()
	if err != nil {
		panic(err)
	}

	productUsecase := NewUseCase(db)

	product, err := productUsecase.GetProduct(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(product.Name)
}

func getDB() (*sql.DB, error) {
	return nil, nil
}
