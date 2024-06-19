package main

import (
	"database/sql"
	"fmt"

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
	p.Price = 1000
	err = updateProduct(db, p)
	if err != nil {
		panic(err)
	}

	product, err := selectProduct(db, p.ID)
	if err != nil {
		panic(err)
	}
	fmt.Print("== Busca por um produto ==\n")
	fmt.Printf("Product: %v, possui o preço de %.2f\n", product.Name, product.Price)

	fmt.Print("== Busca por todos os produtos ==\n")
	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		fmt.Printf("%v, possui o preço de %.2f\n", p.Name, p.Price)
	}

	fmt.Print("== Deletando um produto ==\n")
	err = deleteProduct(db, p.ID)
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

func updateProduct(db *sql.DB, p *Product) error {
	stmt, err := db.Prepare("UPDATE products SET name =?, price =? WHERE id =?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.Name, p.Price, p.ID)
	if err != nil {
		return err
	}
	return nil
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("SELECT id, name, price FROM products WHERE id =?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var p Product
	//err = stmt.QueryRowContext(ctx, id).Scan(&p.ID, &p.Name, &p.Price) É possivel usar o QueryRowContext para passar um contexto com timeout
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM products WHERE id =?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
