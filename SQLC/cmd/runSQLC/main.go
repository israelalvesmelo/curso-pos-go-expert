package main

import (
	"context"
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/israelalvesmelo/sqlc/internal/db"
)

func main() {
	ctx := context.Background()

	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	lastInsertID := createCategory(ctx, queries)

	listCategories(ctx, queries)

	updateCategoty(ctx, lastInsertID, queries)

	listCategories(ctx, queries)

	deleteCategory(ctx, lastInsertID, queries)

	listCategories(ctx, queries)
}

func updateCategoty(ctx context.Context, lastInsertID string, queries *db.Queries) {
	_, err := queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          lastInsertID,
		Name:        "Category Updated",
		Description: sql.NullString{String: "Description Updated", Valid: true},
	})

	if err != nil {
		panic(err)
	}
	println("Category updated")
}

func deleteCategory(ctx context.Context, id string, queries *db.Queries) {
	// Delete Category
	println("Deleting category with id " + id)

	err := queries.DeleteCategory(ctx, id)
	if err != nil {
		panic(err)
	}
	println("Category deleted")
}

func listCategories(ctx context.Context, queries *db.Queries) {
	// List Categories
	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}
}

// TODO: O resultCreated.LastInsertId() retorna um int64 com o valor 0, é necessario ajustar isso.
func createCategory(ctx context.Context, queries *db.Queries) string {
	resultCreated, err := queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:          uuid.New().String(),
		Name:        "Category Created",
		Description: sql.NullString{String: "Description Created", Valid: true},
	})
	if err != nil {
		panic(err)
	}
	lastInsertID, err := resultCreated.LastInsertId()
	if err != nil {
		panic(err)
	}
	id := strconv.FormatInt(lastInsertID, 10)
	println("Category created with id ", id)
	return id
}
