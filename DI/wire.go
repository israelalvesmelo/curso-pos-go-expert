//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/israelalvesmelo/di/product"
)

var setRepository = wire.NewSet(
	product.NewProductRepository,
	wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)),
)

func NewUseCase(db *sql.DB) *product.ProductUsecase {
	wire.Build(
		setRepository,
		product.NewProductUsecase,
	)
	return &product.ProductUsecase{}
}
