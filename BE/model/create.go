package model

import (
	"backend/config"
	"fmt"
)

func Create() {
	// config.Database.Migrator().DropTable(&Order{}, &User{}, &Image{}, &Product{}, &ProductVariance{}, &OrderDetail{}, &User{}, &Page{}, &Collection{}, &Category{}, &CategoryProduct{})
	if err := config.Database.Migrator().AutoMigrate(&Order{},
		&Product{}, &Image{}, &ProductVariance{},
		&OrderDetail{}, &User{}, &Page{}, &Collection{},
		&Category{}, &CategoryProduct{}); err != nil {
		fmt.Print("Error create  table")
	}
}
