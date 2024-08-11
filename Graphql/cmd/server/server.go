package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/israelalvesmelo/graphql/graph"
	"github.com/israelalvesmelo/graphql/graph/generated"
	"github.com/israelalvesmelo/graphql/internal/database"
	_ "github.com/mattn/go-sqlite3"
)

const defaultPort = "8080"

func main() {
	db := setupDatabase()
	defer db.Close()
	categoryDB := database.NewCategory(db)
	courseDB := database.NewCourse(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		CategoryDB: categoryDB,
		CourseDB:   courseDB,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func setupDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	createTableCategories := `CREATE TABLE IF NOT EXISTS categories (
        "id" TEXT,
        "name" TEXT,
        "description" TEXT
    );`

	_, err = db.Exec(createTableCategories)
	if err != nil {
		panic(err)
	}

	createTableCourses := `CREATE TABLE IF NOT EXISTS courses (
        "id" TEXT,
        "name" TEXT,
        "description" TEXT,
		"category_id" TEXT
    );`
	_, err = db.Exec(createTableCourses)
	if err != nil {
		panic(err)
	}

	return db
}
