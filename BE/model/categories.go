package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name             string `json: "name"`
	CollectionID     uint   `json: "collection_id"`
	CategoryProducts []CategoryProduct
}

type CategoryProduct struct {
	gorm.Model
	ProductID  uint `json: "product_id `
	CategoryID uint `json: "category_id"`
}

type Collection struct {
	gorm.Model
	Image      string     `json: "image`
	Name       string     `json: "name"`
	PageID     uint       `json: "page_id"`
	Categories []Category `gorm:"foreignKey:CollectionID"`
}

type Page struct {
	gorm.Model
	Name       string       `json: "name"`
	Collection []Collection `gorm:"foreignKey:PageID"`
}
