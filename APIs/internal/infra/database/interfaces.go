package database

import "github.com/israelalvesmelo/curso-pos-go-expert/APIs/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ProductInterface interface {
	Create(product *entity.Product) error
	FindByID(id int) (*entity.Product, error)
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	Update(product *entity.Product) error
	Delete(id int) error
}