package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Item struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price int
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Item{})

	// Create
	db.Create(&Item{
		Name:  "iPhone",
		Price: 10000,
	})

	// Create multiple records
	// items := []Item{
	// 	{
	// 		Name:  "Samsung",
	// 		Price: 10000,
	// 	},
	// 	{
	// 		Name:  "iPad",
	// 		Price: 20000,
	// 	},
	// }
	// db.Create(&items)

	// select one
	// fmt.Println("=== select one ===")
	// var item Item
	// db.First(&item, 1)
	// fmt.Println(item)
	// var item2 Item

	// db.First(&item2, "name like ?", "iPad")
	// fmt.Println(item2)

	// // select multiple
	// fmt.Println("=== select multiple ===")
	// var items []Item
	// db.Find(&items)
	// for _, item := range items {
	// 	fmt.Println(item)
	// }

	// fmt.Println("=== select with limit and offset ===")
	// var items2 []Item
	// db.Limit(2).Offset(2).Find(&items2)
	// for _, item := range items2 {
	// 	fmt.Println(item)
	// }

	// fmt.Println("=== select with where ===")
	// var items3 []Item
	// db.Where("price >?", 10000).Find(&items3)
	// for _, items3 := range items3 {
	// 	fmt.Println(items3)
	// }

	// // update
	// fmt.Println("=== update ===")
	// var item4 Item
	// db.First(&item4, 1)
	// item4.Name = "iPhone 12"
	// db.Save(&item4)

	var item5 Item
	db.First(&item5, 2)
	fmt.Println(item5)
	// delete
	fmt.Println("=== delete ===")
	db.Delete(&item5)

}
