package repository

import (
	"context"
	"database/sql"

	"github.com/israelalvesmelo/uow/internal/db"
	"github.com/israelalvesmelo/uow/internal/entity"
)

type CategoryRepositoryInterface interface {
	Insert(ctx context.Context, category entity.Category) error
}

type categoryRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewCategoryRepository(dbt *sql.DB) *categoryRepository {
	return &categoryRepository{
		DB:      dbt,
		Queries: db.New(dbt),
	}
}

func (r *categoryRepository) Insert(ctx context.Context, category entity.Category) error {
	return r.Queries.CreateCategory(ctx, db.CreateCategoryParams{
		Name: category.Name,
	})
}
