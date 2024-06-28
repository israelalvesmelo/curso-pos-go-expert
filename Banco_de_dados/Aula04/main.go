package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Items []Item
}

type Item struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Item{}, &Category{})

	// create category
	category := Category{Name: "Cozinha"}
	db.Create(&category)

	// create item
	db.Create(&Item{
		Name:       "Panela",
		Price:      99.0,
		CategoryID: 2,
	})

	var categories []Category
	err = db.Model(&Category{}).Preload("Items").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, item := range category.Items {
			println("- ", item.Name)
		}
	}
}
