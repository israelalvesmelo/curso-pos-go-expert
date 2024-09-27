//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/israelalvesmelo/di/product"
)

func NewUseCase(db *sql.DB) *product.ProductUsecase {
	wire.Build(
		product.NewProductRepository,
		product.NewProductUsecase,
	)
	return &product.ProductUsecase{}
}
