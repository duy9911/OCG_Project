package model

import (
	"backend/dbconfig"
	"fmt"
)

func Create() {
	// dbconfig.Database.Migrator().DropTable(&Order{}, &User{}, &Image{}, &Product{}, &ProductVariance{}, &OrderDetail{}, &User{}, &Page{}, &Collection{}, &Category{}, &CategoryProduct{})
	if err := dbconfig.Database.Migrator().AutoMigrate(&Order{},
		&Product{}, &Image{}, &ProductVariance{},
		&OrderDetail{}, &User{}, &Page{}, &Collection{},
		&Category{}, &CategoryProduct{}); err != nil {
		fmt.Print("Error create  table")
	}
}
