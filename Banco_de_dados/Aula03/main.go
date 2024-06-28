package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Item struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        int
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID     int `gorm:"primaryKey"`
	Number string
	ItemID int
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Item{}, &Category{}, &SerialNumber{})

	// create category
	category := Category{
		Name: "Eletronicos",
	}
	db.Create(&category)

	// create item
	item := Item{
		Name:       "Mouse",
		Price:      1000,
		CategoryID: category.ID,
	}
	db.Create(&item)

	// create serial number
	serialNumber := SerialNumber{
		Number: "123456789",
		ItemID: 1,
	}
	db.Create(&serialNumber)

	var items []Item
	db.Preload("Category").Preload("SerialNumber").Find(&items)
	for _, item := range items {
		fmt.Println(item.Name, item.Category.Name, item.SerialNumber.Number)
	}
}
