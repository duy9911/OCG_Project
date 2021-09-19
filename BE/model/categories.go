package model

import (
	"backend/dbconfig"
	"fmt"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryName     string `json:"category_name"`
	CollectionID     uint   `json:"collection_id"`
	CategoryProducts []CategoryProduct
}

type CategoryProduct struct {
	gorm.Model
	ProductID  uint `json:"product_id"`
	CategoryID uint `json:"category_id"`
}

type Collection struct {
	gorm.Model
	Image          string     `json:"image"`
	CollectionName string     `json:"collection_name"`
	PageID         uint       `json:"page_id"`
	Categories     []Category `gorm:"foreignKey:CollectionID"`
}

type Page struct {
	gorm.Model
	PageName    string       `json:"page_name"`
	Collections []Collection `gorm:"foreignKey:PageID"`
}

// First i create new list as createPage which use to recieve each row from sql  querey returnPagespond
// then i create another struct. i named it OutPut to return value, because one product can be has many properties
// such as size and color
// so enjoy
func OnePageCollections(id string) (Page, error) {
	rt := Page{}
	if err := dbconfig.Database.Where("id = ? ", id).First(&rt).Error; err != nil {
		fmt.Printf(err.Error())
	}
	dbconfig.Database.Model(&rt).Association("Collections").Find(&rt.Collections)
	return rt, nil
}

var pages []Page

//Preload collection first and then match with pages
func AllPageCollections() ([]Page, error) {
	err := dbconfig.Database.Preload("Collections").Find(&pages).Error // SELECT * FROM pages;
	return pages, err                                                  // SELECT * FROM colelctions WHERE colelctions.pages_id IN (1,2,3,4);

}
