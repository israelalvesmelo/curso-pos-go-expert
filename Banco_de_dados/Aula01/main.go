package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float32
}

func NewProduct(name string, price float32) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	p := NewProduct("Macbook", 2000)
	err = insertProduct(db, p)
	if err != nil {
		panic(err)
	}
}

func insertProduct(db *sql.DB, p *Product) error {
	stmt, err := db.Prepare("INSERT INTO products (id, name, price) VALUES (?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.ID, p.Name, p.Price)
	if err != nil {
		return err
	}
	return nil
}
